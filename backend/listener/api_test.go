package listener

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
)

func TestWebhook(t *testing.T) {
	tests := []struct {
		name string
		body figma.FileCommentResponse
	}{
		{
			name: "test valid",
			body: figma.FileCommentResponse{
				Response: figma.Response{
					EventType: "FILE_COMMENT",
					Passcode:  "secretpasscode",
					Timestamp: "2020-02-23T20:27:16Z",
					WebhookID: "22",
				},
				Comment: []figma.CommentFragment{
					{
						Text:    "TODO: \n",
						Mention: "",
					},
					{
						Text:    "",
						Mention: "811724164054158337",
					},
					{
						Text:    "Change selection colors",
						Mention: "",
					},
					{
						Text:    "!issue",
						Mention: "",
					},
				},
				CommentID: "32",
				CreatedAt: "2020-02-23T20:27:16Z",
				FileKey:   "zH44k2FUM629Fa4EMShiHL",
				FileName:  "Mockup library",
				Mentions: []figma.User{
					{
						Id:     "811724164054158337",
						Handle: "Evan Wallace",
						ImgUrl: "",
						Email:  "",
					},
				},
				OrderID:    "23",
				ParentID:   "",
				ResolvedAt: "",
				TriggeredBy: figma.User{
					Id:     "813845097374535682",
					Handle: "Dylan Field",
					ImgUrl: "",
					Email:  "",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			commentChan := make(chan *figma.FileCommentResponse, 8)
			server := Start(context.Background(), commentChan)
			w := httptest.NewRecorder()

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(&tc.body)
			if err != nil {
				t.Fatal(err)
			}
			r, err := http.NewRequest("POST", "/figma", &buf)
			if err != nil {
				t.Fatal(err)
			}
			r.Header.Set("Content-Type", "application/json")
			server.Handler.ServeHTTP(w, r)
		})
	}
}
