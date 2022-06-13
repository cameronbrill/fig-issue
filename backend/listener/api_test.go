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

func TestStart(t *testing.T) {
	tests := []struct {
		name        string
		run         func(t *testing.T, commentChan chan *figma.FigmaFileCommentResponse)
		commentChan chan *figma.FigmaFileCommentResponse
	}{
		{
			name: "test",
			run: func(t *testing.T, commentChan chan *figma.FigmaFileCommentResponse) {
				server := Start(context.Background(), commentChan)
				w := httptest.NewRecorder()
				reqStr := `{
  "comment": [
    {
      "text": "TODO: \n"
    },
    {
      "mention": "811724164054158337"
    },
    {
      "text": "Change selection colors"
    },
    {
      "text": "!issue"
    }
  ],
  "comment_id": "32",
  "created_at": "2020-02-23T20:27:16Z",
  "event_type": "FILE_COMMENT",
  "file_key": "zH44k2FUM629Fa4EMShiHL",
  "file_name": "Mockup library",
  "mentions": [
    {
      "id": "811724164054158337",
      "handle": "Evan Wallace"
    }
  ],
  "order_id": "23",
  "parent_id": "",
  "passcode": "secretpasscode",
  "resolved_at": "",
  "timestamp": "2020-02-23T20:27:16Z",
  "triggered_by": {
    "id": "813845097374535682",
    "handle": "Dylan Field"
  },
  "webhook_id": "22"
 }`
				var req figma.FigmaFileCommentResponse
				err := json.Unmarshal([]byte(reqStr), &req)
				if err != nil {
					t.Fatal(err)
				}

				var buf bytes.Buffer
				err = json.NewEncoder(&buf).Encode(&req)
				if err != nil {
					t.Fatal(err)
				}
				r, err := http.NewRequest("POST", "/figma", &buf)
				if err != nil {
					t.Fatal(err)
				}
				r.Header.Set("Content-Type", "application/json")
				server.Handler.ServeHTTP(w, r)
			},
			commentChan: make(chan *figma.FigmaFileCommentResponse, 8),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.run(t, tc.commentChan)
		})
	}
}
