package words

import (
	"fmt"
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

// TODO: fix this test when the implementation is finalized
func TestGetNewWord(t *testing.T) {
	assert.Equal(t, "hello", getNewWord(time.Now()))
}

func TestOneDaySince(t *testing.T) {
	for i, test := range ts {
		t.Run(fmt.Sprintf("%s-%v", t.Name(), i), func(t *testing.T) {
			assert.Equal(t, test.longer, oneDaySince(test.t1, test.t2))
			assert.Equal(t, test.longer, oneDaySince(test.t2, test.t1))
		})
	}
}
