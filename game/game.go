package game

import (
	"errors"
	"strings"

	"github.com/saxypandabear/wordlego/guess"
)

const DefaultMaxGuesses = 6

type Action string

const (
	// Initiate a new game of Wordle.
	// Acceptable optional inputs:
	// 1. puzzle-num = Solution number for a specific word to guess - defaults to current day
	// 1. max-guesses = configurable maximum number of guesses for the puzzle - defaults to 6
	Start Action = "start"
	// Terminates an active game of Wordle for the player
	Stop Action = "stop"
	// Guesses for the current active game session.
	// Acceptable required inputs:
	// 1. word = the attempted guess for the puzzle
	Guess Action = "guess"
	// Prints out information on the different actions and parameters to the user
	Help Action = "help"
)

// The struct keeps track of an individual user's guesses.
// It also keeps track of the letters individually so it doesn't have
// to be computed over every guess, every time.
// The Letters array uses ternary state. See FormatGuess for the explanation of
// this state. It is duplicated in the Guess struct for simplicity
type WordleSession struct {
	Puzzle            int              // the number of the specific Wordle puzzle
	Solution          string           // the solution for the given session that the player must guess
	Letters           []int            // an array of ints that should be of size 26 to represent the chars
	MessageID         string           // keeps track of the originating message
	Guesses           [][]*guess.Guess // guesses from the user, tracking correctness
	Attempts          []string         // raw guesses from the user
	MaxAllowedGuesses int              // the maximum number of attempts the player has to guess the solution
}

// NewSession creates a new session given a solution, Discord message ID and max number
// of allowed guesses.
func NewSession(solution, messageId string, allowedGuesses, puzzleNum int) *WordleSession {
	ws := WordleSession{
		Puzzle:            puzzleNum,
		Solution:          solution,
		Letters:           make([]int, 26),
		MessageID:         messageId,
		Guesses:           make([][]*guess.Guess, 0, allowedGuesses),
		MaxAllowedGuesses: allowedGuesses,
	}

	return &ws
}

// SetMessageID sets the message ID, because the game session gets instantiated
// prior to the message in the lifecycle. So there has to be a hook in order to
// track the message ID after creation.
func (ws *WordleSession) SetMessageID(messageId string) {
	ws.MessageID = messageId
}

// PrintGame returns a string that represents the state of the session, in order
// to display the game session in Discord chat.
func (ws *WordleSession) PrintGame() string {
	var b strings.Builder
	// TODO: implement
	return b.String()
}

// FormatGuesses takes all of the current guesses in the session, and generates
// the ANSI formatted string to display in Discord that highlights the letters
// in the guesses based on Wordle rules. See wordlego/guess for the formatting
// rules. This function formats all of the guesses with the characters visible.
func (ws *WordleSession) FormatGuesses() string {
	var b strings.Builder
	b.WriteString("```ansi\n") // start ANSI code block
	for _, g := range ws.Guesses {
		b.WriteString(guess.FormatGuess(g) + "\n")
	}
	b.WriteString("```") // close code block
	return b.String()
}

// FormatEmojis takes all of the current guesses in the session, and generates
// the ANSI formatted string to display in Discord that shows all of the guesses
// in emoji form, in the popularized Wordle format.
func (ws *WordleSession) FormatEmojis() string {
	var b strings.Builder
	b.WriteString("```ansi\n")
	for _, g := range ws.Guesses {
		b.WriteString(guess.FormatGuessToEmojis(g) + "\n")
	}
	b.WriteString("```")
	return b.String()
}

// FormatUsedLetters takes all of the Letters and formats a string that illustrates
// the letters that have been used and their correctness.
func (ws *WordleSession) FormatUsedLetters() string {
	// TODO: implement this
	return ""
}

// CanPlay verifies that the number of guesses in the session does not exceed
// the allowed number of guesses for the given session
func (ws *WordleSession) CanPlay() bool {
	return len(ws.Attempts) < ws.MaxAllowedGuesses
}

// Guess attempts to guess with the input word, against the solution, with Worlde
// rules. This assumes that the input exists in the `words/words.go` list of valid
// guesses. This returns an error if the input guess has already been used in this session.
func (ws *WordleSession) Guess(word string) error {
	for _, attempt := range ws.Attempts {
		if attempt == word {
			return errors.New(word + " has already been guessed in this player's session")
		}
	}
	ws.Attempts = append(ws.Attempts, word)
	ws.Guesses = append(ws.Guesses, guess.ConvertToGuess(word, ws.Solution))
	return nil
}
