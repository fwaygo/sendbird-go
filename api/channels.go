package api

type ChannelType string
type InvitationStatus string
type HiddenStatus string

const (
	GROUP_CHANNEL ChannelType = "group_channels"
	OPEN_CHANNEL  ChannelType = "open_channels"
)

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

type ChannelParams struct {
	ChannelType ChannelType
	ChannelUrl  string
}

type ChannelCreateRequest struct {
	UserIDs                 []string                    `json:"user_ids"`
	Name                    *string                     `json:"name,omitempty"`
	ChannelUrl              *string                     `json:"channel_url,omitempty"`
	CoverUrl                *string                     `json:"cover_url,omitempty"`
	CoverFile               []byte                      `json:"cover_file,omitempty"`
	CustomType              *string                     `json:"custom_type,omitempty"`
	IsDistinct              bool                        `json:"is_distinct,omitempty"`
	InvitationStatus        map[string]InvitationStatus `json:"invitation_status,omitempty"`
	Data                    *string                     `json:"data,omitempty"`
	IsPublic                *bool                       `json:"is_public,omitempty"`
	IsEphemeral             *bool                       `json:"is_ephemeral,omitempty"`
	IsSuper                 *bool                       `json:"is_super,omitempty"`
	AccessCode              *string                     `json:"access_code,omitempty"`
	InviterID               *string                     `json:"inviter_id,omitempty"`
	Strict                  *bool                       `json:"strict,omitempty"`
	HiddenStatus            map[string]HiddenStatus     `json:"hidden_status,omitempty"`
	OperatorIDs             []string                    `json:"operator_ids,omitempty"`
	BlockSDKUserChannelJoin *bool                       `json:"block_sdk_user_channel_join,omitempty"`
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
	CreatedAt            int32                   `json:"created_at"`
	Freeze               bool                    `json:"freeze"`
}
