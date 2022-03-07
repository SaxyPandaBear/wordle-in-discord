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
	Puzzle            int            // the number of the specific Wordle puzzle
	Solution          string         // the solution for the given session that the player must guess
	Letters           []int          // an array of ints that should be of size 26 to represent the chars
	MessageID         string         // keeps track of the originating message
	Guesses           []*guess.Guess // guesses from the user, tracking correctness
	Attempts          []string       // raw guesses from the user
	MaxAllowedGuesses int            // the maximum number of attempts the player has to guess the solution
	solved            bool           // flag that is used to determine that the solution has been guessed correctly
}

// NewSession creates a new session given a solution, Discord message ID and max number
// of allowed guesses.
func NewSession(solution, messageId string, allowedGuesses, puzzleNum int) *WordleSession {
	ws := WordleSession{
		Puzzle:            puzzleNum,
		Solution:          solution,
		Letters:           make([]int, 26),
		MessageID:         messageId,
		Guesses:           make([]*guess.Guess, 0, allowedGuesses),
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
// TODO:
// Try to lay it out in an American QWERTY keyboard -
// qwertyuiop
// asdfghjkl
// zxcvbnm
// This is not a requirement to start, but would be good to implement, so leaving the
// todo here.
func (ws *WordleSession) FormatUsedLetters() string {
	// TODO: implement this
	return ""
}

// CanPlay verifies that the number of guesses in the session does not exceed
// the allowed number of guesses for the given session, and the puzzle hasn't
// already been completed
func (ws *WordleSession) CanPlay() bool {
	return !ws.IsSolved() && len(ws.Attempts) < ws.MaxAllowedGuesses
}

func (ws *WordleSession) IsSolved() bool {
	return ws.solved
}

// Guess attempts to guess with the input word, against the solution, with Worlde
// rules. This assumes that the input exists in the `words/wordbank.go` list of valid
// guesses.
// Note that this function DOES have side effects. It updates the state of the game
// session by appending the new guess, updating the colored letters, and updating the
// flag that determines whether or not the solution has been guessed correctly.
// This function returns an error in the scenario where the given word argument
// has already been used in this game session.
func (ws *WordleSession) Guess(word string) error {
	for _, attempt := range ws.Attempts {
		if attempt == word {
			return errors.New(word + " has already been guessed in this player's session")
		}
	}
	if word == ws.Solution {
		ws.solved = true
	}
	ws.Attempts = append(ws.Attempts, word)
	guess := guess.ConvertToGuess(word, ws.Solution)
	ws.updateUsedLetters(guess)
	ws.Guesses = append(ws.Guesses, guess)
	return nil
}

// updateUsedLetters takes a new, valid guess and updates the used letters
// for the game session to reflect the updated correctness of the guess.
// this is simplified because the Guess struct already has all of the used
// letters in the guess itself, and all that needs to be done here is to
// iterate through the used letters and update the values in the Letters array.
// these values are already computed because the int values in the array represent
// the level of correctness, which is already calculated when converting the guessed
// word string into the Guess struct. See guess.ConvertToGuess
// --------------------------------------------------------------------------------
// An open question is whether or not there should be a newly introduced state that
// indicates that a given letter was previously guessed correctly, and then a
// subsequent guess puts the letter in the wrong position.
func (ws *WordleSession) updateUsedLetters(guess *guess.Guess) {
	for _, l := range guess.Letters {
		idx := int(l.Char - 'a') // use 'a' for computing the index
		ws.Letters[idx] = l.Correctness
	}
}
