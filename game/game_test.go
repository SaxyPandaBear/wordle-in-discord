package game

import (
	"strings"
	"testing"

	"github.com/saxypandabear/wordlego/guess"
	"github.com/stretchr/testify/assert"
)

const (
	solution       = "party"
	puzzleNum      = 1
	allowedGuesses = 6
)

func TestCreateNewSession(t *testing.T) {
	ws := NewSession(solution, allowedGuesses, puzzleNum)
	assert.Equal(t, solution, ws.Solution)
	assert.Equal(t, allowedGuesses, ws.MaxAllowedGuesses)
	assert.Equal(t, make([]int, 26), ws.Letters)
	assert.Equal(t, make([]*guess.Guess, 0, allowedGuesses), ws.Guesses)
}

func TestGuess(t *testing.T) {
	ws := testSetup()
	err := ws.Guess("hello")
	assert.NoError(t, err)
	assert.Len(t, ws.Attempts, 1)
	assert.Len(t, ws.Guesses, 1)
	assert.Equal(t, "hello", ws.Attempts[0])
}

func TestGuessRepeatWord(t *testing.T) {
	ws := testSetup()
	g := "hello"
	_ = ws.Guess(g)
	err := ws.Guess(g)
	assert.Error(t, err)
	assert.Len(t, ws.Attempts, 1)
	assert.Equal(t, g, ws.Attempts[0])
	assert.Len(t, ws.Guesses, 1)
	assert.EqualError(t, err, "hello has already been guessed in this player's session")
	err = ws.Guess("beams")
	assert.NoError(t, err)
	assert.Len(t, ws.Attempts, 2)
}

func TestGuessUpdatesLetterCorrectness(t *testing.T) {
	ws := testSetup()
	_ = ws.Guess("pants")

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

func TestSolving(t *testing.T) {
	ws := testSetup()
	assert.True(t, ws.CanPlay())
	assert.False(t, ws.IsSolved())
	err := ws.Guess(solution)
	assert.NoError(t, err)
	assert.True(t, ws.IsSolved())
	// even though there was only one guess for the game session,
	// should not be able to play again because it was solved
	assert.False(t, ws.CanPlay())
}

func TestFormatEmojis(t *testing.T) {
	ws := testSetup()
	ws.Guess("pants")
	ws.Guess("party")
	s := ws.FormatEmojis(true)
	var b strings.Builder
	b.WriteString("```ansi\n")
	b.WriteString(guess.GreenSquare + guess.GreenSquare + guess.BlackSquare + guess.GreenSquare + guess.BlackSquare + "\n")
	b.WriteString(strings.Repeat(guess.GreenSquare, 5) + "\n")
	b.WriteString("```")

	assert.Equal(t, b.String(), s)

	b.Reset()
	s = ws.FormatEmojis(false)
	b.WriteString(guess.GreenSquare + guess.GreenSquare + guess.BlackSquare + guess.GreenSquare + guess.BlackSquare + "\n")
	b.WriteString(strings.Repeat(guess.GreenSquare, 5) + "\n")

	assert.Equal(t, b.String(), s)
}

func testSetup() *WordleSession {
	return NewSession(solution, allowedGuesses, puzzleNum)
}
