package publisher

import (
	"context"

	"github.com/Khan/genqlient/graphql"

	"github.com/cameronbrill/fig-issue/backend/model"
	"github.com/cameronbrill/fig-issue/backend/publisher/linear"
)

type Publisher struct {
	client *graphql.Client
}

func New(gqlc *graphql.Client) *Publisher {
	return &Publisher{client: gqlc}
}

func (p *Publisher) Execute(comment model.Comment) error {
	teams, err := linear.Teams(context.TODO(), *p.client)
	if err != nil {
		return err
	}
	teamId := ""
	for _, t := range teams.Teams.GetNodes() {
		teamId = t.Id
	}
	_, err = linear.IssueCreate(context.TODO(), *p.client, "figma comment", comment.Message, teamId)
	if err != nil {
		return err
	}
	return nil
}
