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

// 性能测试版本1
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStrByRune(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}
	}
}
