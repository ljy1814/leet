package main

import "fmt"

func main() {
	n := 1000
	n = 8765
	n = 9
	n = 20
	n = 99
	fmt.Println(numDupDigitsAtMostN(n))
}

/*
1000
9 9*1
81 9*9
648 9*9*8 最高位不能为0
*/

/*
8765-> [0-1000) [1001-8000) [8001-8700) [8700-8760) [8760-8765)
8780-> 0-1000 1000-8000 8000-8700 8700-8780
*/

func numDupDigitsAtMostN(N int) int {
	if N <= 10 {
		return 0
	}

	cnt := 0
	for t := N; t > 0; {
		cnt++
		t /= 10
	}
	nums := make([]int, cnt+1)
	// N+1 解决20不通过问题
	for t := N + 1; t > 0; {
		nums[cnt] = t % 10
		t /= 10
		cnt--
	}
	if nums[0] == 0 {
		nums = nums[1:]
	}
	cnt = len(nums)
	ret := 0
	var A func(m, n int) int
	A = func(m, n int) int {
		if n == 0 {
			return 1
		}
		return A(m, n-1) * (m - (n - 1))
	}

	// 统计不重复的数字个数 [100...] cnt位
	for i := 1; i < cnt; i++ {
		ret += 9 * A(9, i-1)
	}

	seen := make(map[int]bool)
	for i := 0; i < cnt; i++ {
		j := 1
		if i > 0 {
			j = 0
		}
		for ; j < nums[i]; j++ {
			if !seen[j] {
				t := A(9-i, cnt-i-1)
				ret += t
			}
		}
		if seen[nums[i]] {
			break
		}
		seen[nums[i]] = true
	}
	return N - ret
}
