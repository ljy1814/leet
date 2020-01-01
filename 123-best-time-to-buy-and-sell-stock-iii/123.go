package main

import "fmt"

func maxProfit(prices []int) int {
	helper := func(prices []int, left, right int) int {
		m := 0
		for i := left; i < right; i++ {
			for j := i + 1; j < right; j++ {
				if prices[j]-prices[i] > m {
					m = prices[j] - prices[i]
				}
			}
		}
		return m
	}
	m := 0
	for i := 2; i <= len(prices); i++ {
		m1 := helper(prices, 0, i)
		m2 := helper(prices, i, len(prices))
		if m1+m2 > m {
			m = m1 + m2
		}
	}
	return m
}

func main() {
}

func maxProfit1(prices []int) int {
	pLen := len(prices)
	if pLen == 0 {
		return 0
	}

	result := 0
	profit := make([][3][2]int, pLen)
	/**
	 *三维数组
	 *profit[ii][kk][jj]
	 *ii 第 ii 天， kk 股票操作了几次 ， jj 是否有股票
	 *最多可以完成 两笔 交易： kk 可以为 0 1 2 次操作 ， jj可以为 0 ，1
	 **/
	/**
	 *第一天 初始化
	 *第 1 天 操作 k 次 没有股票，所以初始值为 0
	 *第 1 天 操作i 次 有股票， 所以初始值为 - prices[0]
	 **/
	profit[0][0][0], profit[0][0][1] = 0, -prices[0]
	profit[0][1][0], profit[0][1][1] = 0, -prices[0]
	profit[0][2][0], profit[0][2][1] = 0, -prices[0]

	//注意 买 卖 都进行一次算一次操作 k + 1,单独 买入 不算完成一次操作
	for i := 1; i < pLen; i++ {
		//第 i 天 0 次交易 没有股票最大利润 = 第 i-1 天 0 次交易 没有股票最大利润
		profit[i][0][0] = profit[i-1][0][0]
		//第 i 天 0 次交易 有股票最大利润 = max(第 i-1 天 0 次交易 有股票最大利润 , 第 i-1 天 0 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
		profit[i][0][1] = max(profit[i-1][0][1], profit[i-1][0][0]-prices[i])

		//# 第 i 天 1 次交易 无股票最大利润 = max(第 i-1 天 1次交易 无股票最大利润 , 第 i-1 天 0 次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
		profit[i][1][0] = max(profit[i-1][1][0], profit[i-1][0][1]+prices[i])
		// # 第 i 天 1 次交易 有股票最大利润 = max(第 i-1 天 1 次交易 有股票最大利润 , 第 i-1 天 1 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
		profit[i][1][1] = max(profit[i-1][1][1], profit[i-1][1][0]-prices[i])

		//# 第 i 天 2 次交易 无股票最大利润 = max(第 i-1 天 2次交易 无股票最大利润 , 第 i-1 天 1 次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
		profit[i][2][0] = max(profit[i-1][2][0], profit[i-1][1][1]+prices[i])

	}
	fmt.Println()
	for _, d := range profit {
		fmt.Println(d)
	}
	fmt.Println()
	result = max(profit[pLen-1][0][0], max(profit[pLen-1][1][0], profit[pLen-1][2][0]))
	return result
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func maxProfit2(prices []int) int {
	/*
		// 第i天进行k次交易 0空 1多  利润=本金-购买
		dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i]) // 前一天空状态,今天无操作 前一天多状态,今天空
		dp[i][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1])

		dp[i][0][0] = 0
		dp[i][0][1] = -INF
	*/

	INF := -2 << 31
	k := 2
	dp := [][][2]int{}
	for i := 0; i < len(prices)+1; i++ {
		t := make([][2]int, k+1)
		dp = append(dp, t)
	}

	for _, d := range dp {
		d[0][0] = 0
		d[0][1] = INF
	}

	for i := 1; i <= len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][k][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}

	return dp[len(prices)][k][0]
}

func maxProfit3(prices []int) int {
	/*
		// 第i天进行k次交易 0空 1多  利润=本金-购买
		dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i]) // 前一天空状态,今天无操作 前一天多状态,今天空
		dp[i][k][1] = max(dp[i-1][k-1][0] - prices[i], dp[i-1][k][1])

		dp[i][0][0] = 0
		dp[i][0][1] = -INF
	*/

	type Pair struct {
		Buy  int // 持有股票
		Sell int // 卖出股票
	}

	INF := -2 << 31
	k := 2
	n := len(prices)

	dp := [][]Pair{}
	for i := 0; i <= n; i++ {
		t := make([]Pair, k+1)
		t[0] = Pair{Buy: INF, Sell: 0}
		dp = append(dp, t)
	}

	for j := 1; j <= k; j++ {
		dp[j][0] = Pair{Buy: -prices[j-1], Sell: 0}
		dp[0][j] = Pair{Buy: -prices[j-1], Sell: 0}
	}
	for i := 0; i <= n; i++ {
		fmt.Println(dp[i])
	}

	for i := 1; i <= len(prices); i++ {
		for j := 1; j <= k; j++ {
			p := Pair{}
			fmt.Println(i, j, k, len(dp), len(dp[i-1]), len(prices), dp[i-1][k-1], dp[i-1][k])
			p.Buy = max(dp[i-1][k-1].Sell-prices[i-1], dp[i-1][k].Buy)
			p.Sell = max(dp[i-1][k].Buy+prices[i-1], dp[i-1][k].Sell)
			dp[i][j] = p
		}
	}

	for i := 0; i <= n; i++ {
		fmt.Println(dp[i])
	}

	return dp[n][k].Sell
}

func maxProfit4(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	k := 2
	type Pair struct {
		Buy  int // 持有股票
		Sell int // 卖出股票
	}

	n := len(prices)
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
			//fmt.Println(dp[i-1][j].Buy+prices[i], dp[i-1][j].Sell, dp[i-1][j-1].Sell-prices[i], dp[i-1][j].Buy)
			dp[i][j].Sell = max(dp[i-1][j].Buy+prices[i], dp[i-1][j].Sell)
			dp[i][j].Buy = max(dp[i-1][j-1].Sell-prices[i], dp[i-1][j].Buy)
		}
	}

	return dp[n-1][k].Sell
}
