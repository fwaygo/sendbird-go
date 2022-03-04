package api

import (
	"encoding/json"
	"os"
	"time"
)

type UsersCreateRequest struct {
	UserID           string           `json:"user_id"`
	Nickname         string           `json:"nickname"`
	ProfileURL       string           `json:"profile_url"`
	ProfileFile      *os.File         `json:"profile_file,omitempty"`
	IssueAccessToken *bool            `json:"issue_access_token" default:"false"`
	DiscoveryKeys    []string         `json:"discovery_keys"`
	Metadata         *json.RawMessage `json:"metadata,omitempty"`
}
type UsersCreateResponse = UserResponse

type UsersGetRequest struct {
	UserID             string `json:"user_id"`
	IncludeUnreadCount bool   `json:"include_unread_count" default:"false"`
	CustomTypes        string `json:"custom_types"`
	SuperMode          string `json:"super_mode" default:"all"`
}
type UsersGetResponse = UserResponse

type UsersUpdateRequest struct {
	UserID                  string     `json:"user_id"`
	Nickname                *string    `json:"nickname,omitempty"`
	ProfileURL              *string    `json:"profile_url,omitempty"`
	ProfileFile             *os.File   `json:"profile_file,omitempty"`
	IssueAccessToken        *bool      `json:"issue_access_token,omitempty" default:"false"`
	IsActive                *bool      `json:"is_active,omitempty"`
	LastSeenAt              *time.Time `json:"last_seen_at,omitempty"`
	DiscoveryKeys           []string   `json:"discovery_keys"`
	PreferredLanguages      []string   `json:"preferred_languages"`
	LeaveAllWhenDeactivated *bool      `json:"leave_all_when_deactivated,omitempty" default:"true"`
}
type UsersUpdateResponse = UserResponse

type UsersDeleteRequest struct {
	UserID string `json:"user_id"`
}

type UserResponse struct {
	UserID             string          `json:"user_id"`
	Nickname           string          `json:"nickname"`
	UnreadMessageCount int             `json:"unread_message_count"`
	ProfileURL         string          `json:"profile_url"`
	AccessToken        string          `json:"access_token"`
	IsOnline           bool            `json:"is_online"`
	IsActive           bool            `json:"is_active"`
	CreatedAt          time.Time       `json:"created_at"`
	LastSeenAt         uint64          `json:"last_seen_at"`
	DiscoveryKeys      []string        `json:"discovery_keys"`
	PreferredLanguages []string        `json:"preferred_languages"`
	HasEverLoggedIn    bool            `json:"has_ever_logged_in"`
	Metadata           json.RawMessage `json:"metadata"`
}
