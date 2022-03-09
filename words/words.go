package words

import (
	"fmt"
	"sort"
	"time"
)

var (
	startDate time.Time     = time.Date(2021, time.June, 19, 0, 0, 0, 0, time.UTC) // the reference point for the first Wordle puzzle
	oneDay    time.Duration = time.Hour * 24
)

// WordOfTheDay TBD
func WordOfTheDay(date time.Time) (string, error) {
	if date.Before(startDate) {
		return "", fmt.Errorf("input date %v is invalid because it is before %v", date, startDate)
	}
	return GetSpecificWordleSolution(DetermineWordForDay(date))
}

// GetSpecificWordleSolution is a way to pick a specific problem to play.
// This is expected to align with the numbered problem from Wordle, but this
// accesses solutions by index. As such, we have to subtract 1 from the value
// in order to get the derived index to look up.
// Example:
// To get Wordle 1, the solution is found at solutions[0]
func GetSpecificWordleSolution(num int) (string, error) {
	idx := num - 1
	if idx < 0 || idx >= len(Solutions) {
		return "", fmt.Errorf("input index %d is invalid", idx)
	}
	return Solutions[idx], nil
}

// DetermineWordForDay uses the start date of 06/19/2021 as a reference point
// to derive the index in the solution array. This assumes that the input >= startDate
func DetermineWordForDay(date time.Time) int {
	utc := date.UTC().Truncate(oneDay)
	return int(utc.Sub(startDate).Hours() / oneDay.Hours())
}

// IsGuessValid takes an input string and checks if the string is an allowed guess
// input by checking it against the list of allowed words. This performs a binary search
// for at least some nomimal optimizations
func IsGuessValid(s string) bool {
	sols := GetSortedSolutions()
	idx := sort.Search(len(sols), func(i int) bool {
		return i < len(sols) && sols[i] >= s
	})
	found := idx < len(sols) && sols[idx] == s
	if found {
		return true
	}
	idx = sort.Search(len(AllowedWords), func(i int) bool {
		return i < len(AllowedWords) && AllowedWords[i] >= s
	})
	return idx < len(AllowedWords) && AllowedWords[idx] == s
}
