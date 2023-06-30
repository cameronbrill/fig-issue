package main

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/cameronbrill/fig-issue/backend/model"
	"github.com/cameronbrill/fig-issue/backend/model/figma"
	"github.com/pkg/errors"
)

func transformFigComments(s *figma.FileCommentResponse) (model.Comment, error) {
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

func publishGenericComment(c model.Comment) error {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		return errors.Wrap(err, "creating pulsar client")
	}
	defer client.Close()

	// TODO: all the values below
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-partitioned-topic",
		MessageRouter: func(msg *pulsar.ProducerMessage, tm pulsar.TopicMetadata) int {
			fmt.Println("Topic has", tm.NumPartitions(), "partitions. Routing message ", msg, " to partition 2.")
			// always push msg to partition 2
			return 2
		},
	})
	if err != nil {
		return errors.Wrap(err, "creating pulsar producer")
	}
	defer producer.Close()

	cB, err := json.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "marshalling comment")
	}
	msg := pulsar.ProducerMessage{
		Payload: cB,
		Key:     "message-key",
		Properties: map[string]string{
			"foo": "bar",
		},
		EventTime:           time.Now(),
		ReplicationClusters: []string{"cluster1", "cluster3"},
	}

	if _, err := producer.Send(context.TODO(), &msg); err != nil {
		return errors.Wrap(err, "publishing message to pulsar")
	}

	return nil
}
