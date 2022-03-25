package api

import (
	"encoding/json"
)

type MessageType string
type MentionType string
type ReactionOperation string

const (
	MESG MessageType = "MESG"
	FILE MessageType = "FILE"
	ADMM MessageType = "ADMM"
)

const (
	ADD    ReactionOperation = "ADD"
	DELETE ReactionOperation = "DELETE"
)

type SendBirdFile struct {
	Url  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data string `json:"data"`
}

type BaseMessage struct {
	MessageID       uint64          `json:"message_id"`
	Type            MessageType     `json:"type"`
	CustomType      string          `json:"custom_type"`
	ChannelUrl      string          `json:"channel_url"`
	User            UserResponse    `json:"user"`
	MentionType     MentionType     `json:"mention_type"`
	MentionedUsers  []string        `json:"mentioned_users"`
	IsRemoved       bool            `json:"is_removed"`
	SortedMetaArray json.RawMessage `json:"sorted_metaarray"`
	CreatedAt       uint64          `json:"created_at"`
	UpdatedAt       uint64          `json:"updated_at"`
}

type UserMessage struct {
	BaseMessage
	Message              string          `json:"message"`
	Translations         json.RawMessage `json:"translations"`
	Data                 string          `json:"data"`
	OGTag                json.RawMessage `json:"og_tag"`
	File                 SendBirdFile    `json:"file"`
	IsAppleCriticalAlert bool            `json:"is_apple_critical_alert"`
}

type FileMessage struct {
	BaseMessage
	File        SendBirdFile `json:"file"`
	Thumbnails  []string     `json:"thumbnails"`
	RequireAuth bool         `json:"require_auth"`
}

type Message struct {
	BaseMessage
	Message string          `json:"message"`
	Data    string          `json:"data"`
	OGTag   json.RawMessage `json:"og_tag"`
}

type SendBaseMessageRequest struct {
	MessageType      MessageType     `json:"message_type"`
	CustomType       *string         `json:"custom_type,omitempty"`
	Data             *string         `json:"data,omitempty"`
	SendPush         *bool           `json:"send_push,omitempty"`
	MentionType      *MentionType    `json:"mention_type,omitempty"`
	MentionedUserIDs []string        `json:"mentioned_user_ids,omitempty"`
	IsSilent         *bool           `json:"is_silent,omitempty"`
	SortedMetaArray  json.RawMessage `json:"sorted_metaarray,omitempty"`
	CreatedAt        *uint64         `json:"created_at,omitempty"`
	DedupID          *string         `json:"dedup_id,omitempty"`
	APNSBundleID     *string         `json:"apns_bundle_id,omitempty"`
}

type SendUserMessageRequest struct {
	SendBaseMessageRequest
	UserID                    string          `json:"user_id"`
	Message                   string          `json:"message"`
	AppleCriticalAlertOptions json.RawMessage `json:"apple_critical_alert_options,omitempty"`
	Sound                     *string         `json:"sound,omitempty"`
	Volume                    *float32        `json:"volume,omitempty"`
	MarkAsRead                *bool           `json:"mark_as_read,omitempty"`
}

type SendFileMessageRequest struct {
	SendBaseMessageRequest
	UserID                    string          `json:"user_id"`
	URL                       string          `json:"url"`
	FileName                  *string         `json:"file_name,omitempty"`
	FileSize                  *uint64         `json:"file_size,omitempty"`
	FileType                  *string         `json:"file_type,omitempty"`
	Thumbnails                []string        `json:"thumbnails,omitempty"`
	RequireAuth               *bool           `json:"require_auth,omitempty"`
	SendPush                  *bool           `json:"send_push,omitempty"`
	MarkAsRead                *bool           `json:"mark_as_read,omitempty"`
	AppleCriticalAlertOptions json.RawMessage `json:"apple_critical_alert_options"`
	Sound                     *string         `json:"sound,omitempty"`
	Volume                    *float32        `json:"volume,omitempty"`
}

type AddReactionRequest struct {
	UserID   string `json:"user_id"`
	Reaction string `json:"reaction"`
}

type RemoveReactionRequest struct {
	UserID    string `json:"user_id"`
	MessageID uint64 `json:"message_id"`
	Reaction  string `json:"reaction"`
}

type ReactionUpdateResponse struct {
	UserID    string            `json:"user_id"`
	Operation ReactionOperation `json:"operation"`
	Success   bool              `json:"success"`
	Reaction  string            `json:"reaction"`
	UpdatedAt uint64            `json:"updated_at"`
}
