package main

import "fmt"

func main() {
	str := "cbacdcbc"
	str = "bcabc"
	fmt.Println(removeDuplicateLetters(str))
}

func removeDuplicateLetters(s string) string {
	visited := make([]bool, 26)
	count := make([]int, 26)

	res := "0"
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		count[idx] = count[idx] + 1
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		idx := c - 'a'
		count[idx] = count[idx] - 1
		if visited[idx] {
			continue
		}

		nres := len(res)
		for nres > 0 {
			// 回退  cbacdcbc  cba
			d := res[nres-1]
			ix := d - 'a'
			if c < d && count[ix] > 0 {
				nres--
				visited[ix] = false
				continue
			}
			break
		}
		res = res[:nres]
		res += string(c)
		visited[idx] = true
	}
	return res[1:]
}
