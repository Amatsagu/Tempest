package structs

type ButtonComponentStyle int

const (
	_ ButtonComponentStyle = iota
	PrimaryButtonStyle
	SecondaryButtonStyle
	SuccessButtonStyle
	DangerButtonStyle
	LinkButtonStyle
)

type Message struct {
	Id        string `json:"id"`
	ChannelId string `json:"channel_id"`
	GuildId   string `json:"guild_id,omitempty"`
	TTS       bool   `json:"tts"`
	Pinned    bool   `json:"pinned"`
	/* 	Whether this message mentions everyone */
	MentionEveryone bool `json:"mention_everyone"`
	/* 	Array of users specifically mentioned in the message. */
	Mentions []User `json:"mentions"`
	/* Roles specifically mentioned in this message. */
	MentionRoleIDs  []string `json:"mention_roles"`
	Author          User     `json:"author"`
	Content         string   `json:"content"`
	Timestamp       string   `json:"timestamp,omitempty"`
	EditedTimestamp string   `json:"edited_timestamp,omitempty"`
	Embeds          []Embed  `json:"embeds"`

	// #TODOOOOOOOOOOOOOOO
	Components interface{} `json:"components,omitempty"`
	/* Reference data sent with crossposted messages and inline replies. */
	Reference *MessageReference `json:"message_reference,omitempty"`
	/* ReferencedMessage is the message that was replied to. */
	ReferencedMessage *Message `json:"referenced_message,omitempty"`
}

type MessageReference struct {
	MessageID string `json:"message_id,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
	GuildID   string `json:"guild_id,omitempty"`
}

type MessageButtonComponent struct {
	/* The same as "custom id". */
	Id string `json:"custom_id"`
	/* 	Text label that appears on the button, max 80 characters */
	Text  string               `json:"label,omitempty"`
	Style ButtonComponentStyle `json:"style"`
	/* A url for link-style buttons */
	Url        string `json:"url,omitempty"`
	IsDisabled bool   `json:"disabled,omitempty"`
	/* It gonna always be = 2 for button components. */
	Type int8 `json:"type"`
}
