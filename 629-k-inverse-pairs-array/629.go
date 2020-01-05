package main

import "fmt"

func main() {
	n := 3
	k := 0
	n = 5
	k = 3
	fmt.Println(kInversePairs(n, k))
}

/*
xxxx 5
xxxx5
xxx5x 一定会产生一个逆序对
xx5xx
x5xxx
5xxxx 5放在最前面一定产生4个逆序对
*/

var (
	gdp = [1001][1001]int{}
	M   = 1000000007
)

func dfs(n, k int) int {
	if k < 0 {
		return 0
	}
	if n == 1 {
		if k == 0 {
			return 1
		}
		return 0
	}

	if gdp[n][k] != -1 {
		return gdp[n][k]
	}
	gdp[n][k] = 0

	for i := 1; i <= n; i++ {
		t := i - 1 // n放在第i个位置
		sol := dfs(n-1, k-t)
		gdp[n][k] = (gdp[n][k] + sol) % M
	}
	return gdp[n][k]
}

//dp[5][0]=dp[4][0]
//dp[5][1]=dp[4][1]+dp[4][0]
//dp[5][2]=dp[4][2]+dp[4][1]+dp[4][0]
//dp[5][3]=dp[4][3]+dp[4][2]+dp[4][1]+dp[4][0]
//dp[5][4]=dp[4][4]+dp[4][3]+dp[4][2]+dp[4][1]+dp[4][0]
//dp[5][5]=dp[4][5]+dp[4][4]+dp[4][3]+dp[4][2]+dp[4][1]
//dp[5][6]=dp[4][6]+dp[4][5]+dp[4][4]+dp[4][3]+dp[4][2]

func kInversePairs(n int, k int) int {
	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, k+1)
	}

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		for j := 1; j <= k; j++ {
			// dp[n][k] = dp[n][k-1]+dp[n-1][k] + (k>=n ? -dp[n-1][k-n]:0)
			dp[i][j] = (dp[i-1][j] + dp[i][j-1] + M) % M
			if j >= i {
				dp[i][j] = (dp[i][j] - dp[i-1][j-i] + M) % M
			}
		}
	}
	return dp[n][k]
}
