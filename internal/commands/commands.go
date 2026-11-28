package commands

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type Command struct {
	Name       string
	Definition discord.ApplicationCommandCreate
	Run        func(e *events.ApplicationCommandInteractionCreate) error
}

var All = []Command{
	pingCommand,
	sayCommand,
}

func Definitions() []discord.ApplicationCommandCreate {
	defs := make([]discord.ApplicationCommandCreate, 0, len(All))
	for _, c := range All {
		defs = append(defs, c.Definition)
	}

	return defs
}

func Dispatch(e *events.ApplicationCommandInteractionCreate) {
	name := e.SlashCommandInteractionData().CommandName()
	for _, c := range All {
		if c.Name == name {
			if err := c.Run(e); err != nil {
				fmt.Println("command error: ", err)
			}
			return
		}
	}
}
