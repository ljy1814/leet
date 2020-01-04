package main

import "fmt"

// 1234
// 123 + 1
// 12 * 10 * 10
// 1 * 100 +
// 1 *

// x k x>k x=k x<k
// ? 1 base %  0
func countDigitOne(n int) int {
	k := 1
	base := 1
	ret := 0

	for {
		t := n / base
		if t == 0 {
			break
		}
		ret += t / 10 * base
		m := t % 10
		if m > k {
			ret += base
		} else if m == k {
			ret += n%base + 1
		} else {
			ret += 0
		}
		base *= 10
	}
	return ret
}

func main() {
	n := 1234
	fmt.Println(countDigitOne(n))
}
