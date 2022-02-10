package guess

import "strings"

const (
	GreenSquare  = "🟩"
	YellowSquare = "🟨"
	BlackSquare  = "⬛"
	GreenText    = "[1;32m"
	YellowText   = "[1:33m"
	DefaultText  = "[0m"
)

type Guess struct {
	Letter      rune
	Correctness int
}

func ConvertToGuess(word, solution string) []*Guess {
	guesses := make([]*Guess, len(solution))
	for i, c := range word {
		idx := strings.IndexRune(solution, c)
		var correctness int
		if i == idx {
			correctness = 2
		} else if idx >= 0 {
			correctness = 1
		} else {
			correctness = 0
		}
		guesses[i] = &Guess{
			Letter:      c,
			Correctness: correctness,
		}
	}
	return guesses
}

// FormatGuess takes a word, and returns an ANSI formatted string using the
// correctness level of each rune.
// 0 indicates that the rune is completely invalid ⬛
// 1 indicates that the rune exists in the solution, but is not in the correct position 🟨
// 2 indicates that the rune exists and is the correct position 🟩
// This function returns the runes that are highlighted in their respective colors.
func FormatGuess(guess []*Guess) string {
	var b strings.Builder
	for _, g := range guess {
		var color string
		switch g.Correctness {
		case 0:
			color = DefaultText
		case 1:
			color = YellowText
		case 2:
			color = GreenText
		}
		b.WriteString(color)
		b.WriteRune(g.Letter)
	}
	return b.String()
}

// FormatGuessToEmojis works like FormatGuess, but instead of returning a formatted
// string that displays the guessed runes, this just returns the emoji squares that show
// correctness. This is used to update the original message embed to show progress for the
// latest guess without
func FormatGuessToEmojis(guess []*Guess) string {
	var b strings.Builder
	for _, g := range guess {
		switch g.Correctness {
		case 0:
			b.WriteString(BlackSquare)
		case 1:
			b.WriteString(YellowSquare)
		case 2:
			b.WriteString(GreenSquare)
		}
	}
	return b.String()
}
