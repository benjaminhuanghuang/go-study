/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/
package main

import "fmt"

/*
lastOccurred[x] 不存在 or lastOccurred[x] < start   -> no operation
lastOccurred[x] > start                            -> update start
update lastOccurred[x]
*/
func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLen := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcc"))
	fmt.Println(lengthOfLongestSubstring("存在不存在"))
}
