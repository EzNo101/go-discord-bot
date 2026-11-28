package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var pingCommand = Command{
	Name: "ping",	// for dispatch
	Definition: discord.SlashCommandCreate{
		Name:        "ping",	// for discord
		Description: "Replies with pong",
	},
	Run: func(e *events.ApplicationCommandInteractionCreate) error {
		return e.CreateMessage(discord.NewMessageCreate().WithContent("pong"))
	},
}
