package listener

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
	"github.com/cameronbrill/fig-issue/backend/test"
)

func TestWebhook(t *testing.T) {
	tests := []struct {
		name               string
		body               *figma.FileCommentResponse
		bodyStr            string
		expectedStatusCode int
		expectedStatus     string
		validRequest       bool
	}{
		{
			name:               "test valid",
			expectedStatusCode: http.StatusOK,
			validRequest:       true,
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
		{
			name:               "test invalid invalid passcode",
			expectedStatusCode: http.StatusUnauthorized,
			expectedStatus:     "Invalid passcode",
			validRequest:       true,
			body: &figma.FileCommentResponse{
				Response: figma.Response{
					EventType: "FILE_COMMENT",
					Passcode:  "invalidpasscode",
					Timestamp: "2020-02-23T20:27:16Z",
					WebhookID: "22",
				},
			},
		},
		{
			name:               "test invalid io.ReadCloser provided to request",
			expectedStatusCode: http.StatusBadRequest,
			validRequest:       false,
			body:               nil,
		},
		{
			name:               "test invalid request body",
			expectedStatus:     "Invalid request body",
			expectedStatusCode: http.StatusBadRequest,
			validRequest:       true,
			bodyStr:            `{"comment": 4}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			commentChan := make(chan *figma.FileCommentResponse, 8)
			wbhkSvc := &webhookSvc{commentChan}

			var body io.Reader
			if tc.validRequest {
				var reqBody any = tc.body
				if tc.bodyStr != "" {
					reqBody = bytes.NewBufferString(tc.bodyStr)
				}
				var buf bytes.Buffer
				err := json.NewEncoder(&buf).Encode(reqBody)
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

			if res.StatusCode != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, res.StatusCode)
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(res.Body)
			bodyStr := buf.String()
			if len(tc.expectedStatus) > 0 && !strings.Contains(bodyStr, tc.expectedStatus) {
				t.Errorf("expected status %s, got %s", tc.expectedStatus, bodyStr)
			}
		})
	}
}
