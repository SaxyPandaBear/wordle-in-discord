package game

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/saxypandabear/wordlego/words"
)

// keep track of the active sessions
var sessions = make(map[string]*WordleSession) // TODO: how to make this survive an outage?

type CommandArgs struct {
	GameAction Action
	Word       string
	PuzzleNum  int
	MaxGuesses int
}

// Wordle is the hook for the bot to execute the wordle game functionality.
// This acts as the main game loop.
func Wordle(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	action := data.Options[0].Value.(Action)

	word := ""
	if len(data.Options) >= 2 {
		word = data.Options[1].StringValue()
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
	if _, ok := sessions[i.Message.Author.Username]; !ok {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
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
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "An error occurred when trying to get a solution for your game. Contact the bot owner.",
			},
		})
		return
	}

	gameSession := NewSession(sol, "", args.MaxGuesses, args.PuzzleNum)
	sessions[i.Message.Author.Username] = gameSession

	// TODO: figure this out
	m, err := s.ChannelMessageSend(i.ChannelID, "Wordle "+i.Member.Mention())

	if err != nil {
		delete(sessions, i.Message.Author.Username) // if there was an error, undo the state change
		return
	}
	gameSession.SetMessageID(m.ID)
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
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   1 << 6,
			Content: "Not implemented yet",
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
