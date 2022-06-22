package listener

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
	"github.com/cameronbrill/fig-issue/backend/test"
)

func TestWebhook(t *testing.T) {
	tests := []struct {
		name           string
		body           *figma.FileCommentResponse
		expectedStatus int
		validRequest   bool
	}{
		{
			name:           "test valid",
			expectedStatus: http.StatusOK,
			validRequest:   true,
			body: &figma.FileCommentResponse{
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
			wbhkSvc := &webhookSvc{commentChan}

			var body io.Reader
			if tc.validRequest {
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(tc.body)
				if err != nil {
					t.Fatal(err)
				}
				body = &buf
			} else {
				body = test.ErrReader(0)
			}
			req := httptest.NewRequest(http.MethodPost, "/figma", body)
			w := httptest.NewRecorder()

			wbhkSvc.figmaHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, res.StatusCode)
			}
		})
	}
}
