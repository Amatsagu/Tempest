package structs

import (
	"strings"
	"tempest/misc"
)

type Member struct {
	User            *User    `json:"user"`
	GuildId         string   `json:"-"`
	GuildAvatarHash string   `json:"avatar"`
	Nickname        string   `json:"nickname"`
	JoinedAt        string   `json:"joined_at"`
	NitroSince      string   `json:"premium_since"`
	RoleIds         []string `json:"roles"`
	PermissionFlags int      `json:"permissions" default:"0"`
}

/* Returns a direct url to members's guild specific avatar. It'll return empty string if targeted member don't use custom avatar for that server. */
func (member Member) FetchGuildAvatarUrl() string {
	if member.GuildAvatarHash == "" || member.GuildId == "" {
		return ""
	}

	if strings.HasPrefix(member.GuildAvatarHash, "a_") {
		return misc.DiscordAPIUrl + "/guilds/" + member.GuildId + "/users/" + member.User.Id + "/avatars/" + member.GuildAvatarHash + ".gif"
	}

	return misc.DiscordAPIUrl + "/guilds/" + member.GuildId + "/users/" + member.User.Id + "/avatars/" + member.GuildAvatarHash
}
