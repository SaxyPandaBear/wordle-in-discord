package game

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/saxypandabear/wordlego/words"
)

// keep track of the active sessions
var sessions = make(map[string]*WordleSession) // TODO: how to make this survive an outage?

// Wordle is the hook for the bot to execute the wordle game functionality.
// This acts as the main game loop.
func Wordle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	action := i.ApplicationCommandData().Options[0].StringValue() // this is required, so this is fine
	switch action {
	case "start":
		start(s, i)
	case "stop":
		stop(s, i)
	case "guess":
		guessWord(s, i)
	case "help":
		help(s, i)
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "Invalid action",
			},
		})
	}

	// TODO: Working on this
	solution, err := words.WordOfTheDay(time.Now())
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
		})
		return
	}
	m, err := s.ChannelMessageSendEmbed(i.ChannelID, &discordgo.MessageEmbed{
		Title: "Wordle X 1/6",
		Description: "```ansi\n" +
			`[0;45mABCDEFGHIJK` +
			"\n```",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(m.ID)
	fmt.Println(solution)
}

func start(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Not implemented yet",
		},
	})
}

func stop(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Not implemented yet",
		},
	})
}

func guessWord(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Not implemented yet",
		},
	})
}

// publish a help message to the user
func help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Not implemented yet",
		},
	})
}
