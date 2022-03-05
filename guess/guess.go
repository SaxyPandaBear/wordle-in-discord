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
	Letters []*Letter
}

type Letter struct {
	Char        rune
	Correctness int
}

func (l *Letter) ColoredText() string {
	var col string
	if l.Correctness == 2 {
		col = GreenText
	} else if l.Correctness == 1 {
		col = YellowText
	} else {
		col = DefaultText
	}
	return col + string(l.Char)
}

func (l *Letter) Emoji() string {
	switch l.Correctness {
	case 0:
		return BlackSquare
	case 1:
		return YellowSquare
	case 2:
		return GreenSquare
	default:
		return BlackSquare
	}
}

func ConvertToGuess(word, solution string) *Guess {
	letters := make([]*Letter, len(solution))
	for i, c := range word {
		idx := strings.IndexRune(solution, c)
		var correctness int
		if i == idx {
			correctness = 2
		} else if idx >= 0 {
			correctness = 1
		} else {
			correctness = 0 // TODO: Update this to use -1, and add a 4th color to indicate incorrectness
		}
		letters[i] = &Letter{
			Char:        c,
			Correctness: correctness,
		}
	}
	return &Guess{
		Letters: letters,
	}
}

// FormatGuess takes a word, and returns an ANSI formatted string using the
// correctness level of each rune.
// 0 indicates that the rune is completely invalid ⬛
// 1 indicates that the rune exists in the solution, but is not in the correct position 🟨
// 2 indicates that the rune exists and is the correct position 🟩
// This function returns the runes that are highlighted in their respective colors.
func FormatGuess(guess *Guess) string {
	var b strings.Builder
	for _, l := range guess.Letters {
		b.WriteString(l.ColoredText())
	}
	return b.String()
}

// FormatGuessToEmojis works like FormatGuess, but instead of returning a formatted
// string that displays the guessed runes, this just returns the emoji squares that show
// correctness. This is used to update the original message embed to show progress for the
// latest guess without
func FormatGuessToEmojis(guess *Guess) string {
	var b strings.Builder
	for _, l := range guess.Letters {
		b.WriteString(l.Emoji())
	}
	return b.String()
}
