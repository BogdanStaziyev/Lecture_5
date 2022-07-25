package DoFunc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Do(t *testing.T) {
	tests := map[string]struct {
		s   string
		i   int
		b   bool
		exp string
		err string
	}{
		"successNumber5": {
			"a",
			5,
			true,
			"[5A]",
			"",
		},
		"LowerLetters": {
			"a",
			5,
			false,
			"[5a]",
			"",
		},
		"successNumber13": {
			"a",
			13,
			true,
			"A",
			"",
		},
		"ErrorInvalidNumber": {
			"a",
			9,
			true,
			"",
			"invalid s",
		},
		"ErrorInvalidString": {
			"f",
			5,
			true,
			"",
			"invalid s",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			number, err := Do(tc.s, tc.i, tc.b)
			assert.Equal(t, tc.exp, number)
			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
