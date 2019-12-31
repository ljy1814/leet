package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	dp := make([][]bool, len(p)+1)
	for j := 0; j < len(p)+1; j++ {
		dp[j] = make([]bool, len(s)+1)
	}

	dp[0][0] = true
	for j := 1; j < len(p)+1; j++ {
		if p[j-1] == '*' {
			dp[j][0] = dp[j-1][0]
		}
	}
	// adc
	// a*
	/*
		1 0 0
		1 1 1
	*/

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			m := p[j-1] == '?' || s[i-1] == p[j-1]
			if m {
				dp[j][i] = dp[j-1][i-1]
			}
			if p[j-1] == '*' {
				dp[j][i] = dp[j][i-1] || dp[j-1][i]
			}
		}
	}

	for _, d := range dp {
		for _, t := range d {
			if t == true {
				fmt.Printf("%d ", 1)
			} else {
				fmt.Printf("%d ", 0)
			}
		}
		fmt.Println()
	}
	return dp[len(p)][len(s)]
}

func main() {
	s := "adceb"
	p := "*a*b"
	f := isMatch(s, p)
	fmt.Println(s, p, f)

	s = "adceb"
	p = "a*b"
	f = isMatch(s, p)
	fmt.Println(s, p, f)

	s = "aa"
	p = "a"
	f = isMatch(s, p)
	fmt.Println(s, p, f)

	s = "acdcb"
	p = "a*c?b"
	f = isMatch(s, p)
	fmt.Println(s, p, f)
}
