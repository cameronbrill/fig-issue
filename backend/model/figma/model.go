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

type Response struct {
	EventType WebhookV2Event `json:"event_type" validate:"required"`
	Passcode  string         `json:"passcode" validate:"required"`
	Timestamp string         `json:"timestamp" validate:"required"`
	WebhookID json.Number    `json:"webhook_id" validate:"required"`
}

type PingResponse struct {
	Response
}

type CommentFragment struct {
	Text    string `json:"text"`
	Mention string `json:"mention"`
}

type User struct {
	Id     string `json:"id"`
	Handle string `json:"handle"`
	ImgUrl string `json:"img_url"`
	Email  string `json:"email"`
}

type FileCommentResponse struct {
	Response
	Comment     []CommentFragment `json:"comment"`
	CommentID   json.Number       `json:"comment_id"`
	CreatedAt   string            `json:"created_at"`
	FileKey     string            `json:"file_key"`
	FileName    string            `json:"file_name"`
	Mentions    []User            `json:"mentions"`
	OrderID     json.Number       `json:"order_id,omitempty"`
	ParentID    string            `json:"parent_id"` // int
	ResolvedAt  string            `json:"resolved_at"`
	TriggeredBy User              `json:"triggered_by"`
}
