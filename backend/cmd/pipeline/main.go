package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/Khan/genqlient/graphql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/cameronbrill/fig-issue/backend/listener"
	"github.com/cameronbrill/fig-issue/backend/model"
	"github.com/cameronbrill/fig-issue/backend/model/figma"
	"github.com/cameronbrill/fig-issue/backend/pipeline"
	"github.com/cameronbrill/fig-issue/backend/publisher"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.key)
	return t.wrapped.RoundTrip(req)
}

func transformFigComments(s *figma.FigmaFileCommentResponse) (model.Comment, error) {
	message := ""
	mentions := []string{}

	for _, c := range s.Comment {
		if c.Text == "" {
			continue
		}
		message += c.Text + "\n"
	}
	isIssue, err := regexp.MatchString("!issue", message)
	if err != nil {
		return model.Comment{}, err
	}
	if !isIssue {
		return model.Comment{}, fmt.Errorf("comment is not issue")
	}

	return model.Comment{Message: message, Mentions: mentions}, nil
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %+v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	figCommentChan := make(chan *figma.FigmaFileCommentResponse)
	wbhkSvc := listener.Start(ctx, figCommentChan)
	go func() {
		log.Info("starting figma listener on port :3000")
		err := wbhkSvc.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	lowerStage := make(chan model.Comment)
	errorChannel := make(chan error)

	for i := 0; i < 8; i++ {
		go func() {
			pipeline.Step(ctx, figCommentChan, lowerStage, errorChannel, transformFigComments)
		}()
	}

	key := os.Getenv("LINEAR_API_KEY")
	if key == "" {
		err := fmt.Errorf("must set LINEAR_API_KEY=<linear token>")
		panic(err)
	}
	httpClient := http.Client{
		Transport: &authedTransport{
			key:     key,
			wrapped: http.DefaultTransport,
		},
	}
	client := graphql.NewClient("https://api.linear.app/graphql", &httpClient)
	p := publisher.New(&client)
	pub := func(c model.Comment) error {
		if err := p.Execute(c); err != nil {
			return err
		}
		return nil
	}
	for i := 0; i < 8; i++ {
		err := pipeline.Consumer(ctx, cancel, lowerStage, pub, errorChannel)
		if err != nil {
			log.Fatal(err)
		}
	}

	// TODO: os signal handling
	<-ctx.Done()
}
