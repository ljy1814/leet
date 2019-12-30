package main

import "fmt"

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}

	if s == "" {
		if len(p) == 2 && p[1] == '*' {
			return true
		}
		return false
	}

	if p == "" {
		return false
	}

	firstMatch := false
	if p[0] == '.' || s[0] == p[0] {
		firstMatch = true
	}

	fmt.Println(s, p, firstMatch)
	// aaabc
	// a*bc
	if firstMatch {
		// 匹配多次
		f1 := false
		f2 := false
		if len(p) > 1 && p[1] == '*' {
			f1 = isMatch(s[1:], p)
		} else {
			f2 = isMatch(s[1:], p[1:])
		}
		return f1 || f2
	} else {
		// 匹配0次
		if len(p) > 1 && p[1] == '*' {
			return isMatch(s, p[2:])
		}
	}
	// aab c*a*b
	return firstMatch
}

func isMatch1(s string, p string) bool {
	//dp
	dp := [][]bool{}
	for i := 0; i < len(p)+1; i++ {
		dp = append(dp, make([]bool, len(s)+1))
	}

	dp[0][0] = true
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			dp[i+1][0] = dp[i-1][0]
		}
	}

	//
	// aa
	// a*
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(s); j++ {
			if p[i] == s[j] || p[i] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[i] == '*' {
				ff := p[i-1] == '.' || p[i-1] == s[j]
				dp[i+1][j+1] = dp[i-1][j+1] || ff && dp[i+1][j]
			} else {

			}
		}
	}

	for _, l := range dp {
		fmt.Println(l)
	}
	return dp[len(p)][len(s)]
}
