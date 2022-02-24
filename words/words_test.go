package words

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type times struct {
	t1     time.Time
	t2     time.Time
	longer bool
}

var ts = []times{
	{ // exactly 24 hour difference
		t1:     time.Date(2022, time.February, 9, 1, 2, 3, 4, time.Local),
		t2:     time.Date(2022, time.February, 8, 1, 2, 3, 4, time.Local),
		longer: true,
	},
	{ // <24 hour difference
		t1:     time.Date(2022, time.February, 9, 1, 2, 3, 4, time.Local),
		t2:     time.Date(2022, time.February, 9, 1, 1, 2, 3, time.Local),
		longer: false,
	},
	{ // >24 hour difference because of timezones
		t1:     time.Date(2022, time.February, 9, 18, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60)),
		t2:     time.Date(2022, time.February, 9, 1, 0, 0, 0, time.UTC),
		longer: true,
	},
	{
		t1:     time.Date(2022, time.February, 9, 5, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60)),
		t2:     time.Date(2022, time.February, 9, 5, 0, 0, 0, time.UTC),
		longer: false,
	},
}

func TestIsGuessValidFindsWord(t *testing.T) {
	guess := "abler"
	assert.True(t, IsGuessValid(guess))
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
	assert.Equal(t, solutions[4], actual)
}

func TestGetWordleSolutionInvalidIndex(t *testing.T) {
	_, err := GetSpecificWordleSolution(-1)
	assert.Error(t, err)
	_, err = GetSpecificWordleSolution(len(solutions) + 1) // overflows
	assert.Error(t, err)
}

func TestGetWordleSolutionValidDate(t *testing.T) {
	d := startDate.AddDate(0, 0, 5)
	actual, err := WordOfTheDay(d)
	assert.NoError(t, err)
	assert.Equal(t, solutions[4], actual)
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
