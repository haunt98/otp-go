package hotp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHOTP(t *testing.T) {
	tests := []struct {
		name        string
		key         []byte
		counters    []uint64
		digits      int
		wantResults []int
	}{
		{
			name:     "6 digits",
			key:      []byte("12345678901234567890"),
			counters: []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			digits:   6,
			wantResults: []int{
				755224,
				287082,
				359152,
				969429,
				338314,
				254676,
				287922,
				162583,
				399871,
				520489,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for i := range tc.counters {
				gotResult, gotErr := HOTP(tc.key, tc.counters[i], tc.digits)
				assert.NoError(t, gotErr)
				assert.Equal(t, tc.wantResults[i], gotResult)
			}
		})
	}
}
