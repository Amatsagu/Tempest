package structs

type Embed struct {
	Title        string          `json:"title,omitempty"`
	Url          string          `json:"url,omitempty"`
	Author       *EmbedAuthor    `json:"author,omitempty"`
	Color        uint32          `json:"color,omitempty"`
	ThumbnailUrl *EmbedThumbnail `json:"thumbnail.url,omitempty"`
	Description  string          `json:"description,omitempty"`
	Fields       []*EmbedField   `json:"fields,omitempty"`
	Footer       *EmbedFooter    `json:"footer,omitempty"`
	Image        *EmbedImage     `json:"image.url,omitempty"`
	Video        *EmbedVideo     `json:"video,omitempty"`
	Provider     *EmbedProvider  `json:"provider,omitempty"`
	Timestamp    string          `json:"timestamp,omitempty"`
}

type EmbedAuthor struct {
	IconUrl string `json:"icon_url,omitempty"`
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type EmbedFooter struct {
	IconUrl string `json:"icon_url,omitempty"`
	Text    string `json:"text,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type EmbedVideo struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type EmbedProvider struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}
