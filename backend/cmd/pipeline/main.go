package main

import (
	"context"
	"fmt"
	"regexp"

	log "github.com/sirupsen/logrus"

	"github.com/cameronbrill/fig-issue/backend/listener"
	"github.com/cameronbrill/fig-issue/backend/pipeline"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	figCommentChan := make(chan *listener.FigmaFileCommentResponse)
	go func() {
		listener.Start(ctx, figCommentChan)
	}()

	type comment struct {
		message  string
		mentions []string
	}

	lowerStage := make(chan comment)
	errorChannel := make(chan error)

	transformFigComments := func(s *listener.FigmaFileCommentResponse) (comment, error) {
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
			return comment{}, err
		}
		if !isIssue {
			return comment{}, fmt.Errorf("comment is not issue")
		}

		return comment{message: message, mentions: mentions}, nil
	}

	go func() {
		pipeline.Step(ctx, figCommentChan, lowerStage, errorChannel, transformFigComments)
	}()

	post := func(c comment) error {
		log.Infof("consumed comment: %+#v", c)
		return nil
	}
	err := pipeline.Consumer(ctx, cancel, lowerStage, post, errorChannel)
	if err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
}
