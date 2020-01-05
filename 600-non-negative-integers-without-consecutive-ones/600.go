package main

import "fmt"

func main() {
	n := 13
	fmt.Println(findIntegers(n))
}

/*
0000
0001 1以内1
0010 2以内2
0011
0100 4以内3
0101
0110
0111
1000 8 5
1001
1010
1011
1100
1101
1110
1111
10000 16以内8
f[k]=f[k-1]+f[k-2]
01 10
n=2^k1 + 2^k2 +2^k3 + ...
*/

func findIntegers(num int) int {
	fib := [32]int{1, 2}
	for i := 2; i < 32; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	k := 30
	cons := false
	ret := 1 // 0算
	for k >= 0 {
		if num&(1<<uint(k)) > 0 {
			ret += fib[k]
			if cons {
				// 减去自身
				ret--
				return ret
			}
			cons = true
		} else {
			cons = false
		}
		k--
	}
	return ret
}
