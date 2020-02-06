package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// edge cased
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcabcd", 4},

		// chinses support
		{"我爱慕课网我爱网", 5},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStrByRune(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s: "+
				"expected %d",
				actual, tt.s, tt.ans)
		}
	}
}
