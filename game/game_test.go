package game

import (
	"strings"
	"testing"

	"github.com/saxypandabear/wordlego/guess"
	"github.com/stretchr/testify/assert"
)

const (
	solution       = "party"
	messageId      = "abc123"
	allowedGuesses = 6
)

func TestCreateNewSession(t *testing.T) {
	ws := NewSession(solution, messageId, allowedGuesses)
	assert.Equal(t, solution, ws.Solution)
	assert.Equal(t, messageId, ws.MessageID)
	assert.Equal(t, allowedGuesses, ws.MaxAllowedGuesses)
	assert.Equal(t, make([]int, 26), ws.Letters)
	assert.Equal(t, make([][]*guess.Guess, 0, allowedGuesses), ws.Guesses)
}

func TestSubmitGuess(t *testing.T) {
	ws := testSetup()
	assert.Len(t, ws.Guesses, 0)
	assert.Len(t, ws.Attempts, 0)
	ws.Guess("parts")
	assert.Len(t, ws.Guesses, 1)
	assert.Len(t, ws.Attempts, 1)
}

func TestCanPlay(t *testing.T) {
	ws := testSetup()
	assert.True(t, ws.CanPlay())
	ws.Guess("parts")
	assert.True(t, ws.CanPlay())
	ws.Guess("pools")
	ws.Guess("patty")
	ws.Guess("means")
	ws.Guess("plant")
	// should be able to attempt 1 more
	assert.True(t, ws.CanPlay())
	ws.Guess("heart")
	assert.False(t, ws.CanPlay())
}

func TestDuplicatedGuessDoesNotCountAsNewGuess(t *testing.T) {
	ws := testSetup()
	ws.Guess("parts")
	ws.Guess("parts")
	assert.Len(t, ws.Guesses, 1)
	assert.Len(t, ws.Attempts, 1)
	assert.Equal(t, "parts", ws.Attempts[0])
}

func TestFormatEmojis(t *testing.T) {
	ws := testSetup()
	ws.Guess("pants")
	ws.Guess("party")
	s := ws.FormatEmojis()
	var b strings.Builder
	b.WriteString("```ansi\n")
	b.WriteString(guess.GreenSquare + guess.GreenSquare + guess.BlackSquare + guess.GreenSquare + guess.BlackSquare + "\n")
	b.WriteString(strings.Repeat(guess.GreenSquare, 5) + "\n")
	b.WriteString("```")

	assert.Equal(t, b.String(), s)
}

func testSetup() *WordleSession {
	return NewSession(solution, messageId, allowedGuesses)
}
