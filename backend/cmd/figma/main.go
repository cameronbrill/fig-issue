package main

import (
	"context"

	"github.com/cameronbrill/fig-issue/backend/listener"
	"github.com/cameronbrill/fig-issue/backend/model"
	"github.com/cameronbrill/fig-issue/backend/model/figma"
	"github.com/cameronbrill/fig-issue/backend/pipeline"
	log "github.com/sirupsen/logrus"
)

/*
receive event from figma, preserve if authenticated
filter authenticated events to authenticated comments
transform figma comment to generic comment
post generic comment to pulsar topic
*/
func main() {
	configureLogging()
	log.Info("starting up!")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	figCommentChan := make(chan *figma.FileCommentResponse)
	wbhkSvc := listener.Start(ctx, figCommentChan)
	go func() {
		log.Info("starting figma listener on port :3000")
		err := wbhkSvc.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	genericCommentChan := make(chan model.Comment)
	errorChannel := make(chan error)
	for i := 0; i < 8; i++ {
		go func() {
			pipeline.Step(ctx, figCommentChan, genericCommentChan, errorChannel, transformFigComments)
		}()
	}

	for i := 0; i < 8; i++ {
		go func() {
			pipeline.Consumer(ctx, cancel, genericCommentChan, publishGenericComment, errorChannel)
		}()
	}

	<-ctx.Done()
}
