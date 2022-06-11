package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hasura/go-graphql-client"
	log "github.com/sirupsen/logrus"

	"github.com/cameronbrill/fig-issue/backend/listener"
	"github.com/cameronbrill/fig-issue/backend/model"
	"github.com/cameronbrill/fig-issue/backend/model/figma"
	"github.com/cameronbrill/fig-issue/backend/pipeline"
	"github.com/cameronbrill/fig-issue/backend/publisher"
)

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	figCommentChan := make(chan *figma.FigmaFileCommentResponse)
	for i := 0; i < 8; i++ {
		go func() {
			listener.Start(ctx, figCommentChan)
		}()
	}

	lowerStage := make(chan model.Comment)
	errorChannel := make(chan error)

	for i := 0; i < 8; i++ {
		go func() {
			pipeline.Step(ctx, figCommentChan, lowerStage, errorChannel, transformFigComments)
		}()
	}

	client := graphql.NewClient("https://api.linear.app/graphql")
	p := publisher.New(client)
	pub := func(c model.Comment) error {
		p.Execute(c)
		return nil
	}
	for i := 0; i < 8; i++ {
		err := pipeline.Consumer(ctx, cancel, lowerStage, pub, errorChannel)
		if err != nil {
			log.Fatal(err)
		}
	}

	<-ctx.Done()
}
