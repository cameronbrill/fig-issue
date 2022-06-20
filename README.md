[![Go Report Card](https://goreportcard.com/badge/github.com/cameronbrill/fig-issue/backend)](https://goreportcard.com/report/github.com/cameronbrill/fig-issue/backend)
[![GoDoc](https://godoc.org/github.com/cameronbrill/fig-issue/backend?status.svg)](https://godoc.org/github.com/cameronbrill/fig-issue/backend)

# fig issue

automatically create tickets in your favorite issue platform from figma comments.

## usage

Register on our app; this allows us to register a webhook with your Figma org and to create Linear issues.

## todo

### immediate term

- [x] create base etl pipeline (figma -> fig-issue -> linear)

### long term

- [ ] Add oauth
- [ ] allow extension to other ticket operators (jira, clickup, trello, etc)

## disclaimers

fig issue is not affiliated with Figma. fig issue is simply compatible with Figma. fig issue is not affiliated with Linear.

## thanks

this project uses a lot of open source software. big thanks to the following teams and people for providing them:

- the [mantine team](https://github.com/mantinedev/mantine) for our ui frameworks
  - [rtivital](https://github.com/rtivital) for making lots of example components we use
- [vercel](https://vercel.com) for nextjs and vercel
- [supabase](https://supabase.com) for open source auth
- [Khan Academy](https://github.com/Khan/genqclient) for a modern graphql client for Go
- [Siruspen](github.com/sirupsen/logrus) for structured logging in Go
- [Go Chi](github.com/go-chi/chi) for backend api router
