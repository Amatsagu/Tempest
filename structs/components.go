package structs

type Component interface {
	ButtonComponent | SelectMenuComponent
}

type ButtonComponentStyle int8

const (
	_ ButtonComponentStyle = iota
	PrimaryButtonStyle
	SecondaryButtonStyle
	SuccessButtonStyle
	DangerButtonStyle
	LinkButtonStyle
)

type ButtonComponent struct {
	/* The same as "custom id". */
	Id string `json:"custom_id"`
	/* Text label that appears on the button, max 80 characters. */
	Text  string               `json:"label,omitempty"`
	Emoji *PartialEmoji        `json:"emoji,omitempty"`
	Style ButtonComponentStyle `json:"style"`
	/* A url for link-style buttons */
	Url        string `json:"url,omitempty"`
	IsDisabled bool   `json:"disabled,omitempty"`
	/* It gonna always be = 2 for button components. */
	Type int8 `json:"type"`
}

type SelectMenuComponent struct {
	/* The same as "custom id". */
	Id         string `json:"custom_id"`
	IsDisabled bool   `json:"disabled,omitempty"`
	/* Custom placeholder text if nothing is selected, max 150 characters */
	PlaceholderText string              `json:"placeholder,omitempty"`
	MaxValues       uint32              `json:"max_values,omitempty"`
	MinValues       uint32              `json:"min_values,omitempty"`
	Options         []*SelectMenuOption `json:"options"`
	/* It gonna always be = 3 for select menu components. */
	Type int8 `json:"type"`
}

type SelectMenuOption struct {
	/* Whether to render this option as selected by default. */
	IsDefault bool `json:"default" default:"false"`
	/* Text label that appears on the option label, max 80 characters. */
	Text  string        `json:"label,omitempty"`
	Emoji *PartialEmoji `json:"emoji,omitempty"`
	/* An additional description of the option, max 100 characters. */
	Description string `json:"description,omitempty"`
	/* Dev-defined value of the option, max 100 characters. */
	Value string `json:"value"`
}
