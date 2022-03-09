package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGuessValidFindsWord(t *testing.T) {
	guess := "abler"
	assert.True(t, IsGuessValid(guess))

	guess = "hoard"
	assert.True(t, IsGuessValid(guess)) // word in solutions but not allowed guesses
}

func TestIsGuessValidDoesNotFindWord(t *testing.T) {
	guess := "lllll"
	assert.False(t, IsGuessValid(guess))
}

func TestDetermineWordForDay(t *testing.T) {
	d := startDate.AddDate(0, 0, 2)
	assert.Equal(t, 2, DetermineWordForDay(d))
}

func TestGetWordleSolutionValidIndex(t *testing.T) {
	actual, err := GetSpecificWordleSolution(5)
	assert.NoError(t, err)
	assert.Equal(t, Solutions[4], actual)
}

func TestGetWordleSolutionInvalidIndex(t *testing.T) {
	_, err := GetSpecificWordleSolution(-1)
	assert.Error(t, err)
	_, err = GetSpecificWordleSolution(len(Solutions) + 1) // overflows
	assert.Error(t, err)
}

func TestGetWordleSolutionValidDate(t *testing.T) {
	d := startDate.AddDate(0, 0, 5)
	actual, err := WordOfTheDay(d)
	assert.NoError(t, err)
	assert.Equal(t, Solutions[4], actual)
}

func TestGetWordleSolutionInvalidDate(t *testing.T) {
	d := startDate.AddDate(-1, 0, 0)
	_, err := WordOfTheDay(d)
	assert.Error(t, err)
}

/* benchmark tests for fun */
func BenchmarkIsGuessValidFindsWord(b *testing.B) {
	guesses := []string{ // pick words in different positions of the allowedWords array
		"alure",
		"means",
		"yikes",
	}
	for _, guess := range guesses {
		b.Run(guess, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsGuessValid(guess)
			}
		})
	}
}

func BenchmarkIsGuessValidDoesNotFindWord(b *testing.B) {
	guesses := []string{
		"bcdefg",
		"lmnop",
		"zxywe",
	}
	for _, guess := range guesses {
		b.Run(guess, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsGuessValid(guess)
			}
		})
	}
}
