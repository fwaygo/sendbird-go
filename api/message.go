package api

import (
	"encoding/json"
	"time"
)

type MessageType string
type MentionType string

const (
	MESG MessageType = "MESG"
	FILE MessageType = "FILE"
	ADMM MessageType = "ADMM"
)

type SendBirdFile struct {
	Url  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data string `json:"data"`
}

type BaseMessage struct {
	MessageID       string          `json:"message_id"`
	Type            MessageType     `json:"type"`
	CustomType      string          `json:"custom_type"`
	ChannelUrl      string          `json:"channel_url"`
	User            UserResponse    `json:"user"`
	MentionType     MentionType     `json:"mention_type"`
	MentionedUsers  []string        `json:"mentioned_users"`
	IsRemoved       bool            `json:"is_removed"`
	SortedMetaArray json.RawMessage `json:"sorted_metaarray"`
	CreatedAt       time.Time       `json:"created_at"`
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
