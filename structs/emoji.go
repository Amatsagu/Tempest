package structs

type PartialEmoji struct {
	/* Emoji id may be undefined if that's a default Discord's Tweeter emoji. */
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Animated bool   `json:"animated"`
}
