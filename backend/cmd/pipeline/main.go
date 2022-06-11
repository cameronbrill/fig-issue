package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/cameronbrill/fig-issue/listener"
	"github.com/cameronbrill/fig-issue/pipeline"
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
		fmt.Printf("consumed comment: %+#v\n", c)
		return nil
	}
	err := pipeline.Consumer(ctx, cancel, lowerStage, post, errorChannel)
	if err != nil {
		fmt.Println(err)
	}

	<-ctx.Done()
}
