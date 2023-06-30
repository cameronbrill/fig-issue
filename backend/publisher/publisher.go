package publisher

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

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
	log.Debug("getting linear teams")
	teams, err := linear.Teams(context.TODO(), *p.client)
	if err != nil {
		return errors.Wrap(err, "getting linear teams")
	}
	teamId := ""
	for _, t := range teams.Teams.GetNodes() {
		teamId = t.Id
	}
	log.Debug("creating linear issue")
	_, err = linear.IssueCreate(context.TODO(), *p.client, "figma comment", comment.Message, teamId)
	if err != nil {
		return errors.Wrap(err, "creating linear issue")
	}
	return nil
}
