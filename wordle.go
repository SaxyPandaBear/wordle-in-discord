package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/saxypandabear/wordlego/game"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID  string
	BotToken string
	AppID    string
)

var s *discordgo.Session

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.StringVar(&GuildID, "guild", os.Getenv("GUILDID"), "Test guild ID")
	flag.StringVar(&BotToken, "token", os.Getenv("TOKEN"), "Bot access token")
	flag.StringVar(&AppID, "app", os.Getenv("APPID"), "Application ID")
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

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
	_, err := s.ApplicationCommandCreate(AppID, GuildID, &discordgo.ApplicationCommand{
		Name:        "wordle",
		Description: "Play Wordle! This initiates a new game for the player.",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "action",
				Description: "Action to invoke",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "start",
						Value: game.Start,
					},
					{
						Name:  "stop",
						Value: game.Stop,
					},
					{
						Name:  "guess",
						Value: game.Guess,
					},
					{
						Name:  "help",
						Value: game.Help,
					},
				},
			},
			{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         "word",
				Description:  "Word to guess",
				Required:     false,
				Autocomplete: true,
			},
			{
				Type:         discordgo.ApplicationCommandOptionInteger,
				Name:         "puzzle-num",
				Description:  "Specific puzzle to try to solve. Defaults to the current day",
				Required:     false,
				Autocomplete: true,
			},
			{
				Type:         discordgo.ApplicationCommandOptionInteger,
				Name:         "max-guesses",
				Description:  "Configure the maximum number of guesses for the puzzle",
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
