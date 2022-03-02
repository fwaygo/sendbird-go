package api

import "time"

type InvitationStatus string
type HiddenStatus string

const (
	INVITATION_BY_FRIEND  InvitationStatus = "invited_by_friend"
	INVITED_BY_NON_FRIEND InvitationStatus = "invited_by_non_friend"
	JOINED                InvitationStatus = "joined"
)

const (
	UNHIDDEN                   HiddenStatus = "unhidden"
	HIDDEN_ALLOW_AUTO_UNHIDE   HiddenStatus = "hidden_allow_auto_unhide"
	HIDDEN_PREVENT_AUTO_UNHIDE HiddenStatus = "hidden_prevent_auto_unhide"
)

type ChannelCreateRequest struct {
	UserIDs                 []string                    `json:"user_ids"`
	Name                    string                      `json:"name"`
	ChannelUrl              string                      `json:"channel_url"`
	CoverUrl                string                      `json:"cover_url"`
	CoverFile               []byte                      `json:"cover_file"`
	CustomType              string                      `json:"custom_type"`
	IsDistinct              bool                        `json:"is_distinct"`
	InvitationStatus        map[string]InvitationStatus `json:"invitation_status"`
	Data                    string                      `json:"data"`
	IsPublic                bool                        `json:"is_public"`
	IsEphemeral             bool                        `json:"is_ephemeral"`
	IsSuper                 bool                        `json:"is_super"`
	AccessCode              string                      `json:"access_code"`
	InviterID               string                      `json:"iviter_id"`
	Strict                  bool                        `json:"strict"`
	HiddenStatus            map[string]HiddenStatus     `json:"hidden_status"`
	OperatorIDs             []string                    `json:"operator_ids"`
	BlockSDKUserChannelJoin bool                        `json:"block_sdk_user_channel_join"`
}

type ChannelCreateResponse struct {
	Name                 string                  `json:"name"`
	ChannelUrl           string                  `json:"channel_url"`
	CoverUrl             string                  `json:"cover_url"`
	CustomType           string                  `json:"custom_type"`
	UnreadMessageCount   uint32                  `json:"unread_message_count"`
	Data                 string                  `json:"data"`
	IsDistinct           bool                    `json:"is_distinct"`
	IsPublic             bool                    `json:"is_public"`
	IsSuper              bool                    `json:"is_super"`
	IsEphemeral          bool                    `json:"is_ephemeral"`
	IsAccessCodeRequired bool                    `json:"is_access_code_required"`
	HiddenStatus         map[string]HiddenStatus `json:"hidden_status"`
	MemberCount          uint32                  `json:"member_count"`
	JoinedMemberCount    uint32                  `json:"joined_member_count"`
	Members              []UserResponse          `json:"members"`
	OperatorIDs          []string                `json:"operator_ids"`
	MaxLengthMessage     uint32                  `json:"max_length_message"`
	LastMessage          Message                 `json:"last_message"`
	CreatedAt            time.Time               `json:"created_at"`
	Freeze               bool                    `json:"freeze"`
}
