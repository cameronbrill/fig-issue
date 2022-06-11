package publisher

import (
	"github.com/cameronbrill/fig-issue/backend/model"
	"github.com/hasura/go-graphql-client"
)

type Publisher struct {
	client *graphql.Client
}

func New(gqlc *graphql.Client) *Publisher {
	return &Publisher{client: gqlc}
}

func (p *Publisher) Execute(comment model.Comment) error {
	return nil
}
