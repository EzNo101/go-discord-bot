package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var sayCommand = Command{
	Name: "say",
	Definition: discord.SlashCommandCreate{
		Name:        "say",
		Description: "Makes bot say something.",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "message",
				Description: "What the boy should say",
				Required:    true,
			},
		},
	},
	Run: func(e *events.ApplicationCommandInteractionCreate) error {
		text := e.SlashCommandInteractionData().String("message")

		if err := e.CreateMessage(discord.NewMessageCreate().WithContent("Sent").WithEphemeral(true)); err != nil {
			return err
		}

		_, err := e.Client().Rest.CreateMessage(e.Channel().ID(), discord.NewMessageCreate().WithContent(text))
		return err
	},
}
