package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EzNo101/go-discord-bot/internal/handlers"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/snowflake/v2"
)

func main() {
	// Read the bot token from the DISCORD_TOKEN environment variable.
	// Returns an empty string if the variable is not set.
	token := os.Getenv("DISCORD_TOKEN")
	guildID := snowflake.GetEnv("DISCORD_GUILD_ID")

	// Without a token the bot cannot log in, so stop the program immediately.
	if token == "" {
		panic("DISCORD_TOKEN isn't set")
	}

	// Create the Discord client. disgo.New returns the client and an error.
	// The options after the token configure the client:
	client, err := disgo.New(token,
		// Configure the gateway (the websocket connection to Discord).
		bot.WithGatewayConfigOpts(
			// Only subscribe to guild-related events (servers, channels, roles).
			// Intents tell Discord which events we want to receive.
			gateway.WithIntents(
				gateway.IntentGuilds,
			),
		),
		// Register a handler that runs once the bot has connected and is ready.
		// The function receives an *events.Ready value with info about the bot.
		bot.WithEventListenerFunc(func(e *events.Ready) {
			fmt.Println("Bot is online as", e.User.Username)
		}),
		bot.WithEventListenerFunc(handlers.OnCommand),
	)

	// New can fail (for example, with a bad configuration), so check the error.
	if err != nil {
		panic(err)
	}

	if _, err = client.Rest.SetGuildCommands(client.ApplicationID, guildID, handlers.Commands); err != nil {
		panic(err)
	}

	// defer runs this function when main returns, so the client is always closed.
	// The closure gets its own context with a 10-second timeout for the shutdown.
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// Release the timer resources when this closure finishes.
		defer cancel()
		client.Close(ctx)
	}()

	// Context for opening the connection: it gives up after 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Make sure the timer is cleaned up when main returns.
	defer cancel()

	// Connect to Discord's gateway. The "=" (not ":=") reuses the existing err variable.
	if err = client.OpenGateway(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Bot is running! Press ctrl+c to stop.")

	// Create a channel that can hold one OS signal. A channel is a way for goroutines to pass values.
	stop := make(chan os.Signal, 1)
	// Ask the OS to send us a message on "stop" when the user presses Ctrl+C (SIGINT)
	// or when the process is asked to terminate (SIGTERM).
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// Block here until a signal arrives. Once it does, main returns and the deferred Close runs.
	<-stop
}
