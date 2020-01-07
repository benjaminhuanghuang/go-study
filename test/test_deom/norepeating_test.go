package testdemo

import "testing"

/*
	go test -bench . -cpuprofile cpu.out
	go tool pprof cpu.out
*/

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"阿斯顿发送对方", 6},
		// Edge
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcabcd", 4},
		// Chindese
		{"阿斯顿发送对", 6},
	}

	for _, tt := range tests {	
		if actual := lengthOfLongestSubstring(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}
