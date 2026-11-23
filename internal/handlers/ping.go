package handlers

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var Commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Replies with pong",
	},
}

func OnCommand(e *events.ApplicationCommandInteractionCreate) {
	data := e.SlashCommandInteractionData()
	if data.CommandName() == "ping" {
		err := e.CreateMessage(discord.NewMessageCreate().WithContent("pong"))

		if err != nil {
			fmt.Println("error sending response: ", err)
		}
	}
}
