package main

import "fmt"

func preProcess(s string) string {
	ret := "^"
	for i := 0; i < len(s); i++ {
		ret += "#" + string(s[i])
	}
	ret += "#$"

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome(str string) string {
	println(str)
	T := preProcess(str)
	println(T)
	P := make([]int, len(T))
	C := 0
	R := 0

	for i := 1; i < len(T)-1; i++ {
		iMirror := 2*C - i
		P[i] = 0
		if R > i {
			P[i] = min(R-i, P[iMirror])
		}

		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}
		println("i=", i, "C=", C, "R=", R, "iMirror=", iMirror)
		showI(T)
		fmt.Println(P)
		showT(T)

		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}

	maxLen := 0
	centerIndex := 0

	for i := 1; i < len(T)-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}

	fmt.Println(maxLen, centerIndex)
	return str[(centerIndex-1-maxLen)/2 : maxLen]
}

func showT(T string) {
	print("[")
	for i, _ := range T {
		if i < len(T)-1 {
			fmt.Printf("%c ", T[i])
		} else {
			fmt.Printf("%c", T[i])
		}
	}
	println("]")
}

func showI(T string) {
	print("[")
	for i, _ := range T {
		if i < len(T)-1 {
			fmt.Printf("%d ", i)
		} else {
			fmt.Printf("%d", i)
		}
	}
	println("]")
}

func main() {
	str := "babcbabcbaccba"
	println(longestPalindrome(str))

	//	str = "ABABABA"
	//	println(longestPalindrome(str))
	str = "ac"
	println(longestPalindrome(str))
}
