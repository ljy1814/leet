package main

import "fmt"

/*
 * Author : yajin
 * Email : yajin160305@gmail.com
 * File : reverse.go
 * CreateDate : 2018-12-22 21:10:56
 * */

func reverse(x int) int {

	var (
		digits [10]int
		flag   = false
		i      int
		k      int
		ret    int
	)

	if x < 0 {
		x = -x
		flag = true
	}

	for i, k = 9, x; i >= 0 && k > 0; i-- {
		digits[i] = k % 10
		k /= 10
	}

	k = 9 - i

	// reverse
	fun := func() int {

		for i = 9; i > 9-k; i-- {
			ret = ret*10 + digits[i]
		}

		if flag {
			return -ret
		}
		return ret
	}

	fmt.Println(i, digits)

	if k < 10 {
		return fun()
	}

	if digits[9] > 4 {
		return 0
	}

	if digits[9] == 4 {
		var (
			overflow = 2<<31 - 1
			digits1  [10]int
		)
		// check overflow
		for i = 0; i < 10; i++ {
			digits1[i] = overflow % 10
			overflow /= 10

			if flag && i == 0 {
				digits1[i]++
			}

		}

		for i = 9; i >= 0; i-- {
			if digits[i] > digits1[i] {
				return 0
			}

			if digits[i] < digits1[i] {
				// not overflow
				break
			}
		}

		if i < 0 {
			if flag {
				return -overflow - 1
			}

			return overflow
		}
	}

	return fun()
}

/* vim: set tabstop=4 set shiftwidth=4 */
