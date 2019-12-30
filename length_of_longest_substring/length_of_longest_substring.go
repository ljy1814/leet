package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	window := make(map[byte]int)
	count := 0
	last := -1

	for i := 0; i < len(s); i++ {
		if o, ok := window[s[i]]; ok {
			if o > last {
				last = o
			}
		}
		window[s[i]] = i

		if i-last > count {
			count = i - last
		}
		fmt.Println(i, last, count)
	}

	return count
}

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))

	str = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(str))

	str = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(str))

	str = "au"
	fmt.Println(lengthOfLongestSubstring(str))

	str = "abba"
	fmt.Println(lengthOfLongestSubstring(str))
}
