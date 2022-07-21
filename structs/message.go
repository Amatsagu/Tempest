package structs

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
	Author          *User    `json:"author"`
	Content         string   `json:"content"`
	Timestamp       string   `json:"timestamp,omitempty"`
	EditedTimestamp string   `json:"edited_timestamp,omitempty"`
	Embeds          []*Embed `json:"embeds"`
	/* It should be type of [][]*Component but I don't know how to make Golang happy with struct union... */
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
