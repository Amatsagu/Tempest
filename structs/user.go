package structs

import (
	"strconv"
	"strings"
	"tempest/misc"
)

type Target interface {
	User | Member
}

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	IsBot         bool   `json:"bot" default:"false"`
	/* Hash code used to access user's profile. Call User#FetchAvatarUrl to get direct url. */
	AvatarHash string `json:"avatar"`
	/* Hash code used to access user's baner. Call User#FetchBannerUrl to get direct url. */
	BannerHash  string `json:"banner"`
	PublicFlags int    `json:"public_flags" default:"0"`
	/* User's banner color, encoded as an integer representation of hexadecimal color code. */
	AccentColor int `json:"accent_color"`
	PremiumType int `json:"premium_type"`
}

func (user User) Tag() string {
	return user.Username + "#" + user.Discriminator
}

func (user User) Mention() string {
	return "<@" + user.Username + ">"
}

/* Returns a direct url to user's avatar. It'll return url to default Discord's avatar if targeted user don't use avatar. */
func (user User) FetchAvatarUrl() string {
	if user.AvatarHash == "" {
		id, _ := strconv.Atoi(user.Id)
		return misc.DiscordCDNUrl + "/embed/avatars/" + strconv.Itoa(id%5) + ".png"
	}

	if strings.HasPrefix(user.AvatarHash, "a_") {
		return misc.DiscordCDNUrl + "/avatars/" + user.Id + "/" + user.AvatarHash + ".gif"
	}

	return misc.DiscordCDNUrl + "/avatars/" + user.Id + "/" + user.AvatarHash
}

/* Returns a direct url to user's banner. It'll return empty string if targeted user don't use avatar. */
func (user User) FetchBannerUrl() string {
	if user.BannerHash == "" {
		return ""
	}

	if strings.HasPrefix(user.AvatarHash, "a_") {
		return misc.DiscordCDNUrl + "/banners/" + user.Id + "/" + user.BannerHash + ".gif"
	}

	return misc.DiscordCDNUrl + "/banners/" + user.Id + "/" + user.BannerHash
}
