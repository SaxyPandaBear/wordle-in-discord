package guess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testtable struct {
	input  string
	guess  *Guess
	emojis string
	ansi   string
}

// TODO: Add a test that accounts for duplicate letters
var (
	solution  = "parts"
	testCases = []testtable{
		{
			input: "parts",
			guess: &Guess{
				Letters: []*Letter{
					{
						Char:        'p',
						Correctness: 2,
					},
					{
						Char:        'a',
						Correctness: 2,
					},
					{
						Char:        'r',
						Correctness: 2,
					},
					{
						Char:        't',
						Correctness: 2,
					},
					{
						Char:        's',
						Correctness: 2,
					},
				},
			},
			emojis: GreenSquare + GreenSquare + GreenSquare + GreenSquare + GreenSquare,
			ansi:   fmt.Sprintf("%vp%va%vr%vt%vs", GreenText, GreenText, GreenText, GreenText, GreenText),
		},
		{
			input: "snail",
			guess: &Guess{
				Letters: []*Letter{
					{
						Char:        's',
						Correctness: 1,
					},
					{
						Char:        'n',
						Correctness: 0,
					},
					{
						Char:        'a',
						Correctness: 1,
					},
					{
						Char:        'i',
						Correctness: 0,
					},
					{
						Char:        'l',
						Correctness: 0,
					},
				},
			},
			emojis: YellowSquare + BlackSquare + YellowSquare + BlackSquare + BlackSquare,
			ansi:   fmt.Sprintf("%vs%vn%va%vi%vl", YellowText, DefaultText, YellowText, DefaultText, DefaultText),
		},
		{
			input: "pants",
			guess: &Guess{
				Letters: []*Letter{
					{
						Char:        'p',
						Correctness: 2,
					},
					{
						Char:        'a',
						Correctness: 2,
					},
					{
						Char:        'n',
						Correctness: 0,
					},
					{
						Char:        't',
						Correctness: 2,
					},
					{
						Char:        's',
						Correctness: 2,
					},
				},
			},
			emojis: GreenSquare + GreenSquare + BlackSquare + GreenSquare + GreenSquare,
			ansi:   fmt.Sprintf("%vp%va%vn%vt%vs", GreenText, GreenText, DefaultText, GreenText, GreenText),
		},
	}
)

func TestColoredText(t *testing.T) {
	letter := Letter{
		Char:        'b',
		Correctness: 2,
	}
	assert.Equal(t, GreenText+"b", letter.ColoredText())

	letter.Correctness = 1
	assert.Equal(t, YellowText+"b", letter.ColoredText())

	letter.Correctness = 0
	assert.Equal(t, DefaultText+"b", letter.ColoredText())

	letter.Correctness = -10032431 // random unused number
	assert.Equal(t, DefaultText+"b", letter.ColoredText())
}

func TestEmoji(t *testing.T) {
	letter := Letter{
		Char:        'b',
		Correctness: 2,
	}
	assert.Equal(t, GreenSquare, letter.Emoji())

	letter.Correctness = 1
	assert.Equal(t, YellowSquare, letter.Emoji())

	letter.Correctness = 0
	assert.Equal(t, BlackSquare, letter.Emoji())

	letter.Correctness = 57381234
	assert.Equal(t, BlackSquare, letter.Emoji())
}

func TestConvertToGuess(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			actual := ConvertToGuess(test.input, solution)
			assert.Equal(t, test.guess, actual)
		})
	}
}

func TestFormatGuessToEmojis(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			actual := FormatGuessToEmojis(test.guess)
			assert.Equal(t, test.emojis, actual)
		})
	}
}

func TestFormatGuessToAnsiText(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			actual := FormatGuess(test.guess)
			assert.Equal(t, test.ansi, actual)
		})
	}
}

/* benchmark tests for fun */

func BenchmarkConvertToGuess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToGuess(solution, solution)
	}
}

func BenchmarkConvertGuessToEmojis(b *testing.B) {
	guess := Guess{
		Letters: []*Letter{
			{
				Char:        'p',
				Correctness: 2,
			},
			{
				Char:        'a',
				Correctness: 2,
			},
			{
				Char:        'r',
				Correctness: 2,
			},
			{
				Char:        't',
				Correctness: 2,
			},
			{
				Char:        's',
				Correctness: 2,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		FormatGuessToEmojis(&guess)
	}
}

func BenchmarkConvertGuessToAnsiText(b *testing.B) {
	guess := Guess{
		Letters: []*Letter{
			{
				Char:        'p',
				Correctness: 2,
			},
			{
				Char:        'a',
				Correctness: 2,
			},
			{
				Char:        'r',
				Correctness: 2,
			},
			{
				Char:        't',
				Correctness: 2,
			},
			{
				Char:        's',
				Correctness: 2,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		FormatGuess(&guess)
	}
}
