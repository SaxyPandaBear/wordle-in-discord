package guess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testtable struct {
	input  string
	guess  []*Guess
	emojis string
	ansi   string
}

var (
	solution  = "parts"
	testCases = []testtable{
		{
			input: "parts",
			guess: []*Guess{
				{
					Letter:      'p',
					Correctness: 2,
				},
				{
					Letter:      'a',
					Correctness: 2,
				},
				{
					Letter:      'r',
					Correctness: 2,
				},
				{
					Letter:      't',
					Correctness: 2,
				},
				{
					Letter:      's',
					Correctness: 2,
				},
			},
			emojis: GreenSquare + GreenSquare + GreenSquare + GreenSquare + GreenSquare,
			ansi:   fmt.Sprintf("%vp%va%vr%vt%vs", GreenText, GreenText, GreenText, GreenText, GreenText),
		},
		{
			input: "snail",
			guess: []*Guess{
				{
					Letter:      's',
					Correctness: 1,
				},
				{
					Letter:      'n',
					Correctness: 0,
				},
				{
					Letter:      'a',
					Correctness: 1,
				},
				{
					Letter:      'i',
					Correctness: 0,
				},
				{
					Letter:      'l',
					Correctness: 0,
				},
			},
			emojis: YellowSquare + BlackSquare + YellowSquare + BlackSquare + BlackSquare,
			ansi:   fmt.Sprintf("%vs%vn%va%vi%vl", YellowText, DefaultText, YellowText, DefaultText, DefaultText),
		},
		{
			input: "pants",
			guess: []*Guess{
				{
					Letter:      'p',
					Correctness: 2,
				},
				{
					Letter:      'a',
					Correctness: 2,
				},
				{
					Letter:      'n',
					Correctness: 0,
				},
				{
					Letter:      't',
					Correctness: 2,
				},
				{
					Letter:      's',
					Correctness: 2,
				},
			},
			emojis: GreenSquare + GreenSquare + BlackSquare + GreenSquare + GreenSquare,
			ansi:   fmt.Sprintf("%vp%va%vn%vt%vs", GreenText, GreenText, DefaultText, GreenText, GreenText),
		},
	}
)

func TestConvertToGuess(t *testing.T) {
	for i, test := range testCases {
		t.Run(fmt.Sprintf("%s-%v", t.Name(), i), func(t *testing.T) {
			actual := ConvertToGuess(test.input, solution)
			assert.Equal(t, test.guess, actual)
		})
	}
}

func TestFormatGuessToEmojis(t *testing.T) {
	for i, test := range testCases {
		t.Run(fmt.Sprintf("%s-%v", t.Name(), i), func(t *testing.T) {
			actual := FormatGuessToEmojis(test.guess)
			assert.Equal(t, test.emojis, actual)
		})
	}
}

func TestFormatGuessToAnsiText(t *testing.T) {
	for i, test := range testCases {
		t.Run(fmt.Sprintf("%s-%v", t.Name(), i), func(t *testing.T) {
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
	guess := []*Guess{
		{
			Letter:      'p',
			Correctness: 2,
		},
		{
			Letter:      'a',
			Correctness: 2,
		},
		{
			Letter:      'r',
			Correctness: 2,
		},
		{
			Letter:      't',
			Correctness: 2,
		},
		{
			Letter:      's',
			Correctness: 2,
		},
	}

	for i := 0; i < b.N; i++ {
		FormatGuessToEmojis(guess)
	}
}

func BenchmarkConvertGuessToAnsiText(b *testing.B) {
	guess := []*Guess{
		{
			Letter:      'p',
			Correctness: 2,
		},
		{
			Letter:      'a',
			Correctness: 2,
		},
		{
			Letter:      'r',
			Correctness: 2,
		},
		{
			Letter:      't',
			Correctness: 2,
		},
		{
			Letter:      's',
			Correctness: 2,
		},
	}

	for i := 0; i < b.N; i++ {
		FormatGuess(guess)
	}
}
