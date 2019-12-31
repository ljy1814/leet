package main

import "fmt"

/*
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
S = "ABAACBAB"，T = "ABC"。则问题答案是 "ACB"
*/

func minWindow(s string, t string) string {
	if t == "" || s == "" {
		return s
	}

	cm := make(map[byte]int)
	for _, c := range t {
		cm[byte(c)]++
	}
	index := 0
	offset := len(s) + 1
	hit := 0
	lt := len(t)
	low := 0

	for fast := 0; fast < len(s); fast++ {
		c := byte(s[fast])
		cn := cm[c]
		cm[c]--
		if cn <= 0 {
			continue
		}
		cn--
		// 不在其中
		if cn == 0 {
			hit++
		}
		fmt.Println("----", low, fast, s[low:fast+1], offset, index, cm, hit)
		for hit == lt {
			if fast-low+1 < offset {
				offset = fast - low + 1
				index = low
			}
			fmt.Println(fast, low, s[low:fast+1], offset, index, cm, hit)
			cm[byte(s[low])]++
			x := cm[byte(s[low])]
			low++
			if x > 0 {
				hit--
			}
		}
	}

	fmt.Println(offset, index, cm)
	return s[index : index+offset]
}

func main() {
	S := "ABAACBAB"
	T := "ABC"
	ret := minWindow(S, T)
	fmt.Println(ret)
}
