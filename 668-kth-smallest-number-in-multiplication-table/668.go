package main

func main() {
}

func findKthNumber(m int, n int, k int) int {
	l, r := 1, n*m+1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for l < r {
		mid := l + (r-l)>>1
		cnt := 0
		// 斜三角形,统计小于mid的所有数字
		for i := 1; i <= m; i++ {
			cnt += min(mid/i, n)
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
