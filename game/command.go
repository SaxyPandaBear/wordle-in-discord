package game

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/saxypandabear/wordlego/words"
)

// keep track of the active sessions
var sessions = make(map[string]*WordleSession) // TODO: how to make this survive an outage?

type CommandArgs struct {
	GameAction string
	Word       string
	PuzzleNum  int
	MaxGuesses int
}

// Wordle is the hook for the bot to execute the wordle game functionality.
// This acts as the main game loop.
func Wordle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member == nil {
		// this isn't being called from within a guild. TODO: allow playing Wordle in direct messages
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "Wordle can only be played in a server.",
			},
		})
	}

	args := ParseCommandInputs(i.ApplicationCommandData())
	switch args.GameAction {
	case Start:
		start(s, i, args)
	case Stop:
		stop(s, i, args)
	case Guess:
		guessWord(s, i, args)
	case Help:
		help(s, i, args)
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "Invalid action",
			},
		})
	}
}

// ParseCommandInputs takes the Discord application command data and then
// returns a structured interface that has all of the arguments mapped to
// a comprehensible name, for ease of use. This assumes that the action (at index 0)
// exists, since that is a required argument.
func ParseCommandInputs(data discordgo.ApplicationCommandInteractionData) *CommandArgs {
	action := data.Options[0].Value.(string)

	word := ""
	if len(data.Options) >= 2 {
		word = strings.ToLower(data.Options[1].StringValue())
	}

	var puzzleNum int
	if len(data.Options) >= 3 {
		puzzleNum = int(data.Options[2].IntValue())
	} else {
		puzzleNum = words.DetermineWordForDay(time.Now())
	}

	maxGuesses := DefaultMaxGuesses
	if len(data.Options) >= 4 {
		maxGuesses = int(data.Options[3].IntValue())
	}

	return &CommandArgs{
		GameAction: action,
		Word:       word,
		PuzzleNum:  puzzleNum,
		MaxGuesses: maxGuesses,
	}
}

// start initiates a new game for the user. if the user already has an
// active game session, this emits a failure message to the user indicating such.
func start(s *discordgo.Session, i *discordgo.InteractionCreate, args *CommandArgs) {
	if _, exists := sessions[i.Member.User.ID]; exists {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "You already have an active game. Keep guessing, or use /wordle stop to cancel the active session.",
			},
		})
		return
	}

	sol, err := words.GetSpecificWordleSolution(args.PuzzleNum)
	if err != nil {
		log.Printf("Exception occurred when trying to get a solution for the game: %s\n", err.Error())
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "An error occurred when trying to get a solution for your game. Contact the bot owner.",
			},
		})
		return
	}

	gameSession := NewSession(sol, args.MaxGuesses, args.PuzzleNum)
	sessions[i.Member.User.ID] = gameSession

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: gameSession.PrintGame(false),
		},
	})

	if err != nil {
		delete(sessions, i.Member.User.ID) // if there was an error, undo the state change
		return
	}
}

func stop(s *discordgo.Session, i *discordgo.InteractionCreate, args *CommandArgs) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Not implemented yet",
		},
	})
}

func guessWord(s *discordgo.Session, i *discordgo.InteractionCreate, args *CommandArgs) {
	var sess *WordleSession
	var ok bool
	if sess, ok = sessions[i.Member.User.ID]; !ok {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "You haven't started a game yet. Start one with /wordle start",
			},
		})
		return
	}
	if args.Word == "" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "No guess parameter provided",
			},
		})
		return
	}
	if !words.IsGuessValid(args.Word) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: fmt.Sprintf("'%s' is not a valid guess", args.Word),
			},
		})
		return
	}

	err := sess.Guess(args.Word)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: err.Error(),
			},
		})
		return
	}

	if sess.IsSolved() {
		// player solved the puzzle. share it
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You guessed the word!\n" + sess.PrintGame(true),
			},
		})
		delete(sessions, i.Member.User.ID)
		return
	}
	if !sess.CanPlay() {
		// can't play anymore because the player ran out of tries (different outcome
		// than solving the puzzle).
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You ran out of guesses!\n" + sess.PrintGame(true),
			},
		})
		delete(sessions, i.Member.User.ID)
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: sess.PrintGame(false),
		},
	})
}

// publish a help message to the user
func help(s *discordgo.Session, i *discordgo.InteractionCreate, args *CommandArgs) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Not implemented yet",
		},
	})
}
