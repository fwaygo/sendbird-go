package api

type ChannelListParams string

type ChannelType string
type InvitationStatus string
type HiddenStatus string
type CustomType string
type DistinctMode string
type PublicMode string
type SuperMode string
type Order string

const ()

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

const (
	DISTINCT_ALL DistinctMode = "all"
	DISTINCT     DistinctMode = "distinct"
	NONDISTINCT  DistinctMode = "nondistinct"
)

const (
	PUBLIC_ALL PublicMode = "all"
	PRIVATE    PublicMode = "private"
	PUBLIC     PublicMode = "public"
)

const (
	SUPER_ALL SuperMode = "all"
	SUPER     SuperMode = "super"
	NONSUPER  SuperMode = "nonsuper"
)

const (
	CHRONOLOGICAL               Order = "chronological"
	LATEST_LAST_MESSAGE         Order = "latest_last_message"
	CHANNEL_NAME_ALPHABETICAL   Order = "channel_name_alphabetical"
	METADATA_VALUE_ALPHABETICAL Order = "metadata_value_alphabetical"
)

type ChannelParams struct {
	ChannelType ChannelType
	ChannelUrl  string
	MessageID   *uint64
}

type ChannelCreateRequest struct {
	UserIDs                 []string                    `json:"user_ids"`
	Name                    *string                     `json:"name,omitempty"`
	ChannelUrl              *string                     `json:"channel_url,omitempty"`
	CoverUrl                *string                     `json:"cover_url,omitempty"`
	CoverFile               []byte                      `json:"cover_file,omitempty"`
	CustomType              *CustomType                 `json:"custom_type,omitempty"`
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

type ChannelResponse struct {
	Name                 string                  `json:"name"`
	ChannelUrl           string                  `json:"channel_url"`
	CoverUrl             string                  `json:"cover_url"`
	CustomType           CustomType              `json:"custom_type"`
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

type ChannelListRequest struct {
	Token               *string       `json:"token,omitempty"`
	Limit               *int          `json:"limit,omitempty"`
	DistinctMode        *DistinctMode `json:"distinct_mode,omitempty"`
	PublicMode          *PublicMode   `json:"public_mode,omitempty"`
	SuperMode           *SuperMode    `json:"super_mode,omitempty"`
	CreatedAfter        *uint64       `json:"created_after,omitempty"`
	CreatedBefore       *uint64       `json:"created_before,omitempty"`
	ShowEmpty           bool          `json:"show_empty"`
	ShowMember          bool          `json:"show_member"`
	ShowDeliveryReceipt bool          `json:"show_delivery_receipt"`
	ShowReadReceipt     bool          `json:"show_read_receipt"`
	ShowMetadata        bool          `json:"show_metadata"`
	ShowFrozen          bool          `json:"show_frozen"`
	Order               *Order        `json:"order,omitempty"`
	MetadataOrderKey    *string       `json:"metadata_order_key,omitempty"`
	CustomTypes         *string       `json:"custom_types,omitempty"`
	MembersExactlyIn    []string      `json:"members_exactly_in"`
	// TODO: FINISH THE FIELDS
}

type ChannelListResponse struct {
	Channels []ChannelResponse `json:"channels"`
	Next     *string           `json:"next"`
}

type AddMemberRequest struct {
	ChannelUrl string  `json:"channel_url"`
	UserId     string  `json:"user_id"`
	AccessCode *string `json:"access_code,omitempty"`
}

type AddMemberResponse struct{}

type ChannelHideRequest struct {
	ChannelUrl           string `json:"channel_url"`
	UserID               string `json:"user_id"`
	AllowAutoUnhide      *bool  `json:"allow_auto_unhide,omitempty"`
	ShouldHideAll        *bool  `json:"should_hide_all,omitempty"`
	HidePreviousMessages *bool  `json:"hide_previous_messages,omitempty"`
}

type ChannelUnhideRequest struct {
	ChannelUrl      string `json:"channel_url"`
	UserID          string `json:"user_id"`
	ShouldUnhideAll *bool  `json:"should_unhide_all,omitempty"`
}

type ChannelGetRequest struct {
	ChannelUrl          string `json:"channel_url"`
	ShowDeliveryReceipt *bool  `json:"show_delivery_receipt"`
	ShowReadReceipt     *bool  `json:"show_read_receipt"`
	ShowMember          *bool  `json:"show_member"`
}

type ChannelGetResponse = ChannelResponse
