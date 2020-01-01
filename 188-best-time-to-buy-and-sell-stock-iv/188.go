package main

func main() {
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if k<<1 > n {
		//dp[i][j].Sell = max(dp[i-1][j].Buy+prices[i], dp[i-1][j].Sell)
		//dp[i][j].Buy = max(dp[i-1][j-1].Sell-prices[i], dp[i-1][j].Buy)
		// 当k很大时,可以看成k和k-1近似相同
		// dp[i][0] =max( dp[i-1][0], dp[i-1][0] + prices[i])
		// dp[i][1] =max( dp[i-1][1], dp[i-1][1] - prices[i])
		/*
			dp_i_0 := 0
			dp_i_1 := -2 << 21
			for i := 0; i < n; i++ {
				t := dp_i_0
				dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
				dp_i_1 = max(dp_i_1, t-prices[i])
			}
		*/
		// 交易次数超过n/2可认为交易不限制
		p := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				p += prices[i] - prices[i-1]
			}
		}
		return p
	}

	type Pair struct {
		Buy  int // 持有股票
		Sell int // 卖出股票
	}

	dp := make([][]Pair, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]Pair, k+1)
	}
	// 第一天买入股票,不论买几次
	for i := 0; i <= k; i++ {
		dp[0][i].Buy = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j].Sell = max(dp[i-1][j].Buy+prices[i], dp[i-1][j].Sell)
			dp[i][j].Buy = max(dp[i-1][j-1].Sell-prices[i], dp[i-1][j].Buy)
		}
	}

	return dp[n-1][k].Sell
}
