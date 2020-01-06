package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := "13"
	n = "14549"
	n = "99834"
	fmt.Println(n, nearestPalindromic(n))
}

func nearestPalindromic(n string) string {
	// 99999 100001 对称
	// 14549 -> 14541 14441 14641
	// 99834 -> 99799 99899 99999
	// 764323 -> 763367 764467 765567
	l := len(n)
	num, _ := strconv.ParseInt(n, 10, 64)
	if l <= 1 {
		if n == "0" {
			return "1"
		}
		return strconv.FormatInt(num-1, 10)
	}
	tmpRes := []int64{}
	var k int64 = 1
	for i := 1; i < l; i++ {
		k *= 10
	}
	tmpRes = append(tmpRes, k-1)
	tmpRes = append(tmpRes, k*10+1)

	reverse := func(s string) string {
		t := ""
		for i := len(s) - 1; i >= 0; i-- {
			t += string(s[i])
		}
		return t
	}

	for i := -1; i < 2; i++ {
		half := ""
		if l%2 == 0 {
			half = n[:l>>1]
		} else {
			half = n[:(l+1)>>1]
		}
		tn, _ := strconv.ParseInt(half, 10, 64)
		tn = tn + int64(i)
		half = strconv.FormatInt(tn, 10)

		if l%2 == 0 {
			half += reverse(half)
		} else {
			half += reverse(half[:len(half)-1])
		}

		tn, _ = strconv.ParseInt(half, 10, 64)
		tmpRes = append(tmpRes, tn)
	}

	abs := func(a int64) int64 {
		if a < 0 {
			return -a
		}
		return a
	}
	k = 2 << 31
	var ret int64 = 0
	fmt.Println(tmpRes)
	for _, t := range tmpRes {
		t1 := abs(t - num)
		if t1 == 0 {
			continue
		}
		if t1 < k {
			k = t1
			ret = t
		}
	}

	return strconv.FormatInt(ret, 10)
}
