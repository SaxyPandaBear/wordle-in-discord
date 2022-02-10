package words

import (
	"fmt"
	"sync"
)

var (
	lock sync.Mutex // synchronizes picking the word of the day
)

func WordOfTheDay(day int) string {
	lock.Lock()
	defer lock.Unlock()

	var word string
	if day%2 == 0 {
		word = "hello"
	} else {
		word = "olleh"
	}
	fmt.Printf("Word of the day is: %s\n", word)
	return word
}
