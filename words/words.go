package words

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const hoursInADay float64 = 24.0

var (
	lock           sync.Mutex // synchronizes picking the word of the day
	dtWordAcquired *time.Time
	currentWord    string
	oneDay         time.Duration = time.Hour * 24
)

// WordOfTheDay returns the singleton value of currentWord. This is synchronized with a lock
// in order to avoid a race condition where multiple users could request a new word for the day
// and the word gets overwritten. For the same date, expected to be in UTC time, it should return
// the same word. When given a word
func WordOfTheDay(currentTime time.Time) string {
	lock.Lock()
	defer lock.Unlock()

	if dtWordAcquired == nil || oneDaySince(currentTime, *dtWordAcquired) {
		currentWord = getNewWord(currentTime)
		dtWordAcquired = &currentTime
	}

	fmt.Printf("Word of the day is: %s\n", currentWord)
	return currentWord
}

// getNewWord TBD
func getNewWord(ts time.Time) string {
	return "hello"
}

// oneDaySince checks two time structs, and returns true if the delta
// is at least 24 hours. it returns false otherwise. This function forcefully
// truncates the time arguments, and converts them to UTC time for consistency
func oneDaySince(currentTime, pastTime time.Time) bool {
	t1 := currentTime.UTC().Truncate(oneDay)
	t2 := pastTime.UTC().Truncate(oneDay)

	return math.Abs(float64(t1.Sub(t2).Hours())) >= hoursInADay
}
