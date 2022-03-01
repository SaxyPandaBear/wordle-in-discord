package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/saxypandabear/wordlego/game"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID  = flag.String("guild", os.Getenv("GUILD_ID"), "Test guild ID")
	BotToken = flag.String("token", os.Getenv("TOKEN"), "Bot access token")
	AppID    = flag.String("app", os.Getenv("APP_ID"), "Application ID")
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

// Important note: call every command in order it's placed in the example.
var (
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"wordle": game.Wordle,
	}
)

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is up!")
	})
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		}
	})

	// Wordle game command registration
	// This command is a single entrypoint for the Wordle game.
	// It accepts different subcommand options (via a choice of set values),
	// along with all of the necessary subcommand options for each of those actions.
	_, err := s.ApplicationCommandCreate(*AppID, *GuildID, &discordgo.ApplicationCommand{
		Name:        "wordle",
		Description: "Play Wordle! This initiates a new game for the player.",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "action",
				Description: "Action to invoke",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "start",
						Value: "start",
					},
					{
						Name:  "stop",
						Value: "stop",
					},
					{
						Name:  "guess",
						Value: "guess",
					},
					{
						Name:  "help",
						Value: "help",
					},
				},
			},
			{
				Type:         discordgo.ApplicationCommandOptionInteger,
				Name:         "puzzle-num",
				Description:  "Specific puzzle to try to solve. Defaults to the current day",
				Required:     false,
				Autocomplete: true,
			},
		},
	})

	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}
