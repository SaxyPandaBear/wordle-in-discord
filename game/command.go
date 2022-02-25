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
