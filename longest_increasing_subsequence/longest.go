package main

import "fmt"

func longest(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}

	dp := make([]int, len(arr))
	dp[0] = 1
	for i := 1; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	fmt.Println(dp)
	return dp[len(arr)-1]
}

func alllongest(arr []int) {
	n := len(arr)
	if n <= 0 {
		return
	}

	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	ret := [][]int{}
	kk := dp[n-1]
	var dfs func(m, n, k int, res []int)
	dfs = func(m, n, k int, res []int) {
		if k == 0 {
			ret = append(ret, res)
			return
		}
		for j := n - 1; j >= 0; j-- {
			if dp[j] == k && arr[j] < m {
				t := append(res, j)
				dfs(arr[j], j, k-1, t)
			}
		}
		return
	}

	for j := n - 1; j >= 0; j-- {
		if dp[j] == kk {
			dfs(arr[j], j, kk-1, []int{j})
		}
	}
	fmt.Println(ret)
	return
}

func main() {
	arr := []int{
		10, 9, 2, 5, 3, 8, 7, 101, 18,
	}
	fmt.Println(longest(arr))

	alllongest(arr)
}
