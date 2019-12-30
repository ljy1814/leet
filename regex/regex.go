package main

import "fmt"

/*
a?b**bcd
*/

func cleanPatten(pat string) string {
	if pat == "" || len(pat) == 1 {
		return pat
	}

	var pre, cur byte
	pre = pat[0]
	ret := []byte{pre}
	for i := 1; i < len(pat); i++ {
		cur = pat[i]
		if pre == '*' && cur == '*' {
			pre = cur
			continue
		}

		pre = cur
		ret = append(ret, cur)
	}

	return string(ret)
}

func regex(text, pat string, debug bool) bool {
	if text == "" || pat == "" {
		return false
	}

	fmt.Println("TEXT", text, "PAT", pat)
	i, j := 0, 0
	for i < len(text) {
		if j >= len(pat) {
			break
		}
		if debug {
			fmt.Printf("text[%d]=%c,pat[%d]=%c\n", i, text[i], j, pat[j])
		}
		if pat[j] == '?' || text[i] == pat[j] {
			i++
			j++
			continue
		}

		if pat[j] == '*' && j < len(pat)-1 {
			return regex(text[i:], pat[j+1:], debug)
		}
		if j > 0 {
		} else {
			i++
		}
		j = 0
	}

	return j >= len(pat)-1
}

func main() {
	pat := "abc?"
	fmt.Println(pat, cleanPatten(pat))
	text := "helloababcdefg"
	fmt.Println(pat, text, regex(text, pat, false))

	pat = "abc?*?b"
	fmt.Println(pat, cleanPatten(pat))
	text = "helloababcdefgbbb"
	fmt.Println(pat, text, regex(text, pat, false))

	pat = "abc?*bd"
	fmt.Println(pat, cleanPatten(pat))
	text = "helloababcdefgbbbcbfghbdbcbde"
	fmt.Println(pat, text, regex(text, pat, false))

	pat = "abc?*bd*cdf"
	fmt.Println(pat, cleanPatten(pat))
	text = "helloababcdefgbbbcbfghbdbcbdecdhlscdfg"
	fmt.Println(pat, text, regex(text, pat, true))

	pat = "abc*?*?b"
	fmt.Println(pat, cleanPatten(pat))

	pat = "abc****?*?b"
	fmt.Println(pat, cleanPatten(pat))
}
