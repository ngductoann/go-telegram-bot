package types

// TelegramUpdate represents an update from Telegram
type TelegramUpdate struct {
	UpdateID          int64                      `json:"update_id"`
	Message           *TelegramMessage           `json:"message,omitempty"`
	EditedMessage     *TelegramMessage           `json:"edited_message,omitempty"`
	ChannelPost       *TelegramMessage           `json:"channel_post,omitempty"`
	EditedChannelPost *TelegramMessage           `json:"edited_channel_post,omitempty"`
	InlineQuery       *TelegramInlineQuery       `json:"inline_query,omitempty"`
	CallbackQuery     *TelegramCallbackQuery     `json:"callback_query,omitempty"`
	ShippingQuery     *TelegramShippingQuery     `json:"shipping_query,omitempty"`
	PreCheckoutQuery  *TelegramPreCheckoutQuery  `json:"pre_checkout_query,omitempty"`
	Poll              *TelegramPoll              `json:"poll,omitempty"`
	PollAnswer        *TelegramPollAnswer        `json:"poll_answer,omitempty"`
	MyChatMember      *TelegramChatMemberUpdated `json:"my_chat_member,omitempty"`
	ChatMember        *TelegramChatMemberUpdated `json:"chat_member,omitempty"`
	ChatJoinRequest   *TelegramChatJoinRequest   `json:"chat_join_request,omitempty"`
}

type TelegramInlineQuery struct {
	ID       string        `json:"id"`
	From     *TelegramUser `json:"from"`
	Query    string        `json:"query"`
	Offset   string        `json:"offset"`
	ChatType *string       `json:"chat_type,omitempty"`
}

type TelegramCallbackQuery struct {
	ID              string           `json:"id"`
	From            *TelegramUser    `json:"from"`
	Message         *TelegramMessage `json:"message,omitempty"`
	InlineMessageID *string          `json:"inline_message_id,omitempty"`
	ChatInstance    string           `json:"chat_instance"`
	Data            *string          `json:"data,omitempty"`
	GameShortName   *string          `json:"game_short_name,omitempty"`
}

type TelegramChatMemberUpdated struct {
	Chat          *TelegramChat           `json:"chat"`
	From          *TelegramUser           `json:"from"`
	Date          int64                   `json:"date"`
	OldChatMember *TelegramChatMember     `json:"old_chat_member"`
	NewChatMember *TelegramChatMember     `json:"new_chat_member"`
	InviteLink    *TelegramChatInviteLink `json:"invite_link,omitempty"`
}

type TelegramChatMember struct {
	User   *TelegramUser `json:"user"`
	Status string        `json:"status"`
}

type TelegramChatInviteLink struct {
	InviteLink  string        `json:"invite_link"`
	Creator     *TelegramUser `json:"creator"`
	IsPrimary   bool          `json:"is_primary"`
	IsRevoked   bool          `json:"is_revoked"`
	Name        *string       `json:"name,omitempty"`
	ExpireDate  *int64        `json:"expire_date,omitempty"`
	MemberLimit *int          `json:"member_limit,omitempty"`
}

type TelegramPollAnswer struct {
	PollID    string        `json:"poll_id"`
	User      *TelegramUser `json:"user"`
	OptionIDs []int         `json:"option_ids"`
}

type TelegramChatJoinRequest struct {
	Chat       *TelegramChat           `json:"chat"`
	From       *TelegramUser           `json:"from"`
	Date       int64                   `json:"date"`
	Bio        *string                 `json:"bio,omitempty"`
	InviteLink *TelegramChatInviteLink `json:"invite_link,omitempty"`
}

type TelegramShippingQuery struct {
	ID              string           `json:"id"`
	From            *TelegramUser    `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *TelegramAddress `json:"shipping_address"`
}

type TelegramAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type TelegramPreCheckoutQuery struct {
	ID               string             `json:"id"`
	From             *TelegramUser      `json:"from"`
	Currency         string             `json:"currency"`
	TotalAmount      int                `json:"total_amount"`
	InvoicePayload   string             `json:"invoice_payload"`
	ShippingOptionID *string            `json:"shipping_option_id,omitempty"`
	OrderInfo        *TelegramOrderInfo `json:"order_info,omitempty"`
}

type TelegramOrderInfo struct {
	Name            *string          `json:"name,omitempty"`
	PhoneNumber     *string          `json:"phone_number,omitempty"`
	Email           *string          `json:"email,omitempty"`
	ShippingAddress *TelegramAddress `json:"shipping_address,omitempty"`
}

// TelegramMessage represents a Telegram message object
type TelegramMessage struct {
	// Basic fields
	MessageID       int64   `json:"message_id"`
	MessageThreadID *int64  `json:"message_thread_id,omitempty"`
	Date            int64   `json:"date"`
	Text            *string `json:"text,omitempty"`

	// Sender and chat information
	From *TelegramUser `json:"from,omitempty"`
	Chat *TelegramChat `json:"chat,omitempty"`

	// Forwarding information
	ForwardFrom      *TelegramUser `json:"forward_from,omitempty"`
	ForwardFromChat  *TelegramChat `json:"forward_from_chat,omitempty"`
	ForwardFromMsgID *int64        `json:"forward_from_message_id,omitempty"`
	ForwardDate      *int64        `json:"forward_date,omitempty"`

	// Reply information
	ReplyToMessage *TelegramMessage `json:"reply_to_message,omitempty"`

	// Media attachments
	Photo     []*TelegramPhotoSize `json:"photo,omitempty"`
	Audio     *TelegramAudio       `json:"audio,omitempty"`
	Document  *TelegramDocument    `json:"document,omitempty"`
	Video     *TelegramVideo       `json:"video,omitempty"`
	Sticker   *TelegramSticker     `json:"sticker,omitempty"`
	Voice     *TelegramVoice       `json:"voice,omitempty"`
	Animation *TelegramAnimation   `json:"animation,omitempty"`

	// Other content types
	Location *TelegramLocation `json:"location,omitempty"`
	Contact  *TelegramContact  `json:"contact,omitempty"`
	Poll     *TelegramPoll     `json:"poll,omitempty"`

	// Service messages
	NewChatMembers     []*TelegramUser      `json:"new_chat_members,omitempty"`
	LeftChatMember     *TelegramUser        `json:"left_chat_member,omitempty"`
	NewChatTitle       *string              `json:"new_chat_title,omitempty"`
	NewChatPhoto       []*TelegramPhotoSize `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto    *bool                `json:"delete_chat_photo,omitempty"`
	GroupChatCreated   *bool                `json:"group_chat_created,omitempty"`
	SupergroupCreated  *bool                `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated *bool                `json:"channel_chat_created,omitempty"`
}

// TelegramUser represents a Telegram user or bot
type TelegramUser struct {
	ID           TelegramUserID `json:"id"`
	IsBot        bool           `json:"is_bot"`
	FirstName    string         `json:"first_name"`
	LastName     *string        `json:"last_name,omitempty"`
	Username     *string        `json:"username,omitempty"`
	LanguageCode *string        `json:"language_code,omitempty"`
	IsPremium    *bool          `json:"is_premium,omitempty"`
}

// TelegramChat represents a Telegram chat
type TelegramChat struct {
	ID        TelegramChatID `json:"id"`
	Type      ChatType       `json:"type"`
	Title     *string        `json:"title,omitempty"`
	Username  *string        `json:"username,omitempty"`
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
}

// TelegramPhotoSize represents a size of a photo or a file/sticker thumbnail
type TelegramPhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize *int   `json:"file_size,omitempty"`
}

// TelegramDocument represents a general file
type TelegramDocument struct {
	FileID   string  `json:"file_id"`
	FileName *string `json:"file_name,omitempty"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *int    `json:"file_size,omitempty"`
}

// TelegramVideo represents a video file
type TelegramVideo struct {
	FileID   string  `json:"file_id"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	Duration int     `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *int    `json:"file_size,omitempty"`
}

// TelegramSticker represents a sticker
type TelegramSticker struct {
	FileID string `json:"file_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// TelegramVoice represents a voice note
type TelegramVoice struct {
	FileID   string  `json:"file_id"`
	Duration int     `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *int    `json:"file_size,omitempty"`
}

// TelegramAudio represents an audio file
type TelegramAudio struct {
	FileID   string  `json:"file_id"`
	Duration int     `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize *int    `json:"file_size,omitempty"`
}

// TelegramAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound)
type TelegramAnimation struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`
}

// TelegramLocation represents a point on the map
type TelegramLocation struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// TelegramContact represents a phone contact
type TelegramContact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
}

// TelegramPoll represents a poll
type TelegramPoll struct {
	ID                    string        `json:"id"`
	Question              string        `json:"question"`
	Options               []*PollOption `json:"options"`
	TotalVoterCount       int           `json:"total_voter_count"`
	IsClosed              bool          `json:"is_closed"`
	IsAnonymous           bool          `json:"is_anonymous"`
	Type                  string        `json:"type"`
	AllowsMultipleAnswers bool          `json:"allows_multiple_answers"`
	CorrectOptionID       *int          `json:"correct_option_id,omitempty"`
	Explanation           *string       `json:"explanation,omitempty"`
}

// PollOption represents an option in a poll
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}
