// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package linear

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// IssueCreateIssueCreateIssuePayload includes the requested fields of the GraphQL type IssuePayload.
type IssueCreateIssueCreateIssuePayload struct {
	// Whether the operation was successful.
	Success bool `json:"success"`
	// The issue that was created or updated.
	Issue IssueCreateIssueCreateIssuePayloadIssue `json:"issue"`
}

// GetSuccess returns IssueCreateIssueCreateIssuePayload.Success, and is useful for accessing the field via an interface.
func (v *IssueCreateIssueCreateIssuePayload) GetSuccess() bool { return v.Success }

// GetIssue returns IssueCreateIssueCreateIssuePayload.Issue, and is useful for accessing the field via an interface.
func (v *IssueCreateIssueCreateIssuePayload) GetIssue() IssueCreateIssueCreateIssuePayloadIssue {
	return v.Issue
}

// IssueCreateIssueCreateIssuePayloadIssue includes the requested fields of the GraphQL type Issue.
// The GraphQL type's documentation follows.
//
// An issue.
type IssueCreateIssueCreateIssuePayloadIssue struct {
	// The unique identifier of the entity.
	Id string `json:"id"`
	// The issue's title.
	Title string `json:"title"`
}

// GetId returns IssueCreateIssueCreateIssuePayloadIssue.Id, and is useful for accessing the field via an interface.
func (v *IssueCreateIssueCreateIssuePayloadIssue) GetId() string { return v.Id }

// GetTitle returns IssueCreateIssueCreateIssuePayloadIssue.Title, and is useful for accessing the field via an interface.
func (v *IssueCreateIssueCreateIssuePayloadIssue) GetTitle() string { return v.Title }

// IssueCreateResponse is returned by IssueCreate on success.
type IssueCreateResponse struct {
	// Creates a new issue.
	IssueCreate IssueCreateIssueCreateIssuePayload `json:"issueCreate"`
}

// GetIssueCreate returns IssueCreateResponse.IssueCreate, and is useful for accessing the field via an interface.
func (v *IssueCreateResponse) GetIssueCreate() IssueCreateIssueCreateIssuePayload {
	return v.IssueCreate
}

// TeamsResponse is returned by Teams on success.
type TeamsResponse struct {
	// All teams whose issues can be accessed by the user. This might be different from `administrableTeams`, which also includes teams whose settings can be changed by the user.
	Teams TeamsTeamsTeamConnection `json:"teams"`
}

// GetTeams returns TeamsResponse.Teams, and is useful for accessing the field via an interface.
func (v *TeamsResponse) GetTeams() TeamsTeamsTeamConnection { return v.Teams }

// TeamsTeamsTeamConnection includes the requested fields of the GraphQL type TeamConnection.
type TeamsTeamsTeamConnection struct {
	Nodes []TeamsTeamsTeamConnectionNodesTeam `json:"nodes"`
}

// GetNodes returns TeamsTeamsTeamConnection.Nodes, and is useful for accessing the field via an interface.
func (v *TeamsTeamsTeamConnection) GetNodes() []TeamsTeamsTeamConnectionNodesTeam { return v.Nodes }

// TeamsTeamsTeamConnectionNodesTeam includes the requested fields of the GraphQL type Team.
// The GraphQL type's documentation follows.
//
// An organizational unit that contains issues.
type TeamsTeamsTeamConnectionNodesTeam struct {
	// The unique identifier of the entity.
	Id string `json:"id"`
	// The team's name.
	Name string `json:"name"`
}

// GetId returns TeamsTeamsTeamConnectionNodesTeam.Id, and is useful for accessing the field via an interface.
func (v *TeamsTeamsTeamConnectionNodesTeam) GetId() string { return v.Id }

// GetName returns TeamsTeamsTeamConnectionNodesTeam.Name, and is useful for accessing the field via an interface.
func (v *TeamsTeamsTeamConnectionNodesTeam) GetName() string { return v.Name }

// __IssueCreateInput is used internally by genqlient
type __IssueCreateInput struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	TeamId string `json:"teamId"`
}

// GetTitle returns __IssueCreateInput.Title, and is useful for accessing the field via an interface.
func (v *__IssueCreateInput) GetTitle() string { return v.Title }

// GetBody returns __IssueCreateInput.Body, and is useful for accessing the field via an interface.
func (v *__IssueCreateInput) GetBody() string { return v.Body }

// GetTeamId returns __IssueCreateInput.TeamId, and is useful for accessing the field via an interface.
func (v *__IssueCreateInput) GetTeamId() string { return v.TeamId }

func IssueCreate(
	ctx context.Context,
	client graphql.Client,
	title string,
	body string,
	teamId string,
) (*IssueCreateResponse, error) {
	__input := __IssueCreateInput{
		Title:  title,
		Body:   body,
		TeamId: teamId,
	}
	var err error

	var retval IssueCreateResponse
	err = client.MakeRequest(
		ctx,
		"IssueCreate",
		`
mutation IssueCreate ($title: String!, $body: String!, $teamId: String!) {
	issueCreate(input: {title:$title,description:$body,teamId:$teamId}) {
		success
		issue {
			id
			title
		}
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func Teams(
	ctx context.Context,
	client graphql.Client,
) (*TeamsResponse, error) {
	var err error

	var retval TeamsResponse
	err = client.MakeRequest(
		ctx,
		"Teams",
		`
query Teams {
	teams {
		nodes {
			id
			name
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}