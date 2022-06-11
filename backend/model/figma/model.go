package figma

import "encoding/json"

type WebhookV2Event string

const (
	WebhookV2EventPing        = WebhookV2Event("PING")
	WebhookV2EventFileComment = WebhookV2Event("FILE_COMMENT")
)

type WebhookPostRequest struct {
	EventType WebhookV2Event `json:"event_type"`
}

type FigmaResponse struct {
	EventType WebhookV2Event `json:"event_type"`
	Passcode  string         `json:"passcode"`
	Timestamp string         `json:"timestamp"`
	WebhookID json.Number    `json:"webhook_id"`
}

type FigmaPingResponse struct {
	FigmaResponse
}

type FigmaCommentFragment struct {
	Text    string `json:"text"`
	Mention string `json:"mention"`
}

type FigmaUser struct {
	Id     string `json:"id"`
	Handle string `json:"handle"`
	ImgUrl string `json:"img_url"`
	Email  string `json:"email"`
}

type FigmaFileCommentResponse struct {
	FigmaResponse
	Comment     []FigmaCommentFragment `json:"comment"`
	CommentID   json.Number            `json:"comment_id"`
	CreatedAt   string                 `json:"created_at"`
	FileKey     string                 `json:"file_key"`
	FileName    string                 `json:"file_name"`
	Mentions    []FigmaUser            `json:"mentions"`
	OrderID     json.Number            `json:"order_id,omitempty"`
	ParentID    string                 `json:"parent_id"` // int
	ResolvedAt  string                 `json:"resolved_at"`
	TriggeredBy FigmaUser              `json:"triggered_by"`
}
