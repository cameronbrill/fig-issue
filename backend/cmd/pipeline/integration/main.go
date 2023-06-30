package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/cameronbrill/fig-issue/backend/model/figma"
	"github.com/pkg/errors"
)

func main() {
	mockFigmaRes := figma.FileCommentResponse{
		Response: figma.Response{
			EventType: "FILE_COMMENT",
			Passcode:  "secretpasscode",
			Timestamp: "123",
			WebhookID: "",
		},
		Comment: []figma.CommentFragment{
			{Text: "!issue\n\nThis is a test issue"},
		},
		CommentID:  "",
		CreatedAt:  "",
		FileKey:    "",
		FileName:   "",
		Mentions:   []figma.User{},
		OrderID:    "",
		ParentID:   "",
		ResolvedAt: "",
		TriggeredBy: figma.User{
			Id:     "",
			Handle: "",
			ImgUrl: "",
			Email:  "",
		},
	}
	mockFigmaResB, err := json.Marshal(mockFigmaRes)
	if err != nil {
		panic(errors.Wrap(err, "marshalling mock figma response"))
	}
	mockFigmaResBuf := bytes.NewBuffer(mockFigmaResB)
	const endpoint = "http://localhost:3000/figma"
	res, err := http.Post(endpoint, "application/json", mockFigmaResBuf)
	if err != nil {
		panic(errors.Wrap(err, "making request to webhook"))
	}
	defer res.Body.Close()
}
