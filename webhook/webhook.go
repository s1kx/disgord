package webhook

import (
	"github.com/andersfylling/disgord/user"

	"github.com/andersfylling/snowflake"
)

type Webhook struct {
	ID        snowflake.ID `json:"id"`
	GuildID   snowflake.ID `json:"guild_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	User      *user.User   `json:"user"`
	Name      string       `json:"name"`
	Avatar    string       `json:"avatar"`
	Token     string       `json:"token"`
}
