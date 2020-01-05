package main

import (
	"fmt"
)

func main() {
	n := 9
	n = 15
	fmt.Println(consecutiveNumbersSum(n))
}

/*
15
15 * 1 +0
7 + 8 = 7*2 + 1
4+5+6 = 4 * 3 + 1 + 2
1+2+3+4+5 = 1*5 + 1 + 2 + 3 + 4
*/
func consecutiveNumbersSum(N int) int {
	if N <= 0 {
		return 0
	}
	ret := 0
	for i := 1; N > 0; N, i = N-i, i+1 {
		if (N % i) == 0 {
			ret += 1
		}
	}
	return ret
}
