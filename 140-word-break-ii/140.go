/*
 * Author : yajin
 * Email : yajin160305@gmail.com
 * File : 140.go
 * CreateDate : 2019-12-27 23:45:54
 * */

package main

import (
	"fmt"
	"strings"
)

/*
输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
*/

func wordBreak1(s string, wordDict []string) []string {
	return wordBreak_1(s, wordDict, 0)
}

func wordBreak_1(s string, wordDict []string, start int) []string {
	res := []string{}
	if start == len(s) {
		res = append(res, "")
	}

	for end := start + 1; end <= len(s); end++ {
		contain := false
		for _, word := range wordDict {
			if word == s[start:end] {
				contain = true
				break
			}
		}

		if contain {
			tmp := wordBreak_1(s, wordDict, end)
			for _, l := range tmp {
				space := ""
				if l != "" {
					space = " "
				}
				res = append(res, s[start:end]+space+l)
			}
		}
	}

	return res
}

var (
	bufmap = make(map[int][]string)
)

func wordBreak(s string, wordDict []string) []string {
	return wordBreak_2(s, wordDict, 0)
}

func wordBreak_2(s string, wordDict []string, start int) []string {
	if bufmap[start] != nil {
		return bufmap[start]
	}
	res := []string{}
	if start == len(s) {
		res = append(res, "")
	}

	for end := start + 1; end <= len(s); end++ {
		contain := false
		for _, word := range wordDict {
			if word == s[start:end] {
				contain = true
				break
			}
		}

		if contain {
			tmp := wordBreak_2(s, wordDict, end)
			for _, l := range tmp {
				space := ""
				if l != "" {
					space = " "
				}
				res = append(res, s[start:end]+space+l)
			}
		}
	}

	return res
}

func wordBreak0(s string, wordDict []string) []string {
	dp := make([][]string, len(s)+1)
	initial := []string{""}
	dp[0] = initial

	for i := 1; i <= len(s); i++ {
		l := []string{}
		for j := 0; j < i; j++ {
			contain := false
			for _, w := range wordDict {
				if s[j:i] == w {
					contain = true
				}
			}

			if len(dp[j]) > 0 && contain {
				for _, d := range dp[j] {
					space := ""
					if d != "" {
						space = " "
					}
					l = append(l, d+space+s[j:i])
				}
			}
		}
		dp[i] = l
	}
	return dp[len(s)]
}

func count(wordDict []string, s string) int {
	c := 0
	for _, w := range wordDict {
		if w == s {
			c++
		}
	}

	return c
}

type TrieNode struct {
	flag bool
	next map[byte]*TrieNode
}

func initTrie(wordDict []string) *TrieNode {
	root := &TrieNode{}

	for _, w := range wordDict {
		n := root
		for _, c := range w {
			if n.next == nil {
				n.next = make(map[byte]*TrieNode)
			}
			if n.next[byte(c)] == nil {
				n.next[byte(c)] = &TrieNode{}
			}
			n = n.next[byte(c)]
		}
		n.flag = true
	}

	return root
}

func dfs(s string, dp [][]int, i int, v []string, res []string) []string {
	if i == 0 {
		t := ""
		for k := 0; k < len(v); k++ {
			t += v[len(v)-1-k]
			if k != len(v)-1 {
				t += " "
			}
		}
		res = append(res, t)
		return res
	}

	fmt.Println(i, dp[i])
	for _, j := range dp[i] {
		v = append(v, s[j:i])
		fmt.Println(i, j, v)
		res = dfs(s, dp, j, v, res)
		v = v[:len(v)-1]
		fmt.Println(i, j, v, "---")
	}

	return res
}

func wordBreak9(s string, wordDict []string) []string {
	root := initTrie(wordDict)
	n := len(s)
	dp := make([][]int, n+1)
	dp[0] = []int{}
	dp[0] = append(dp[0], -1)
	for i := 0; i < n; i++ {
		if dp[i] == nil {
			continue
		}
		t := root
		for j := i; j < n && t.next[s[j]] != nil; {
			t = t.next[s[j]]
			j++
			if t.flag {
				if dp[j] == nil {
					dp[j] = []int{}
				}
				dp[j] = append(dp[j], i)
			}
		}
	}
	fmt.Println(dp)

	v := []string{}
	res := []string{}
	res = dfs(s, dp, n, v, res)
	return res
}

func wordBreak8(s string, wordDict []string) []string {
	wMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wMap[v] = true
	}
	//fmt.Println(wMap)
	//先通过DP判断字符串是否能被拆分
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)

	re := []string{}
	if !dp[len(s)] {
		return re
	}
	//回溯走起
	var DFS = func(string, []string) {}
	DFS = func(ns string, r []string) {
		if len(ns) == 0 {
			re = append(re, strings.Join(r, " "))
			return
		}
		for i := 1; i <= len(ns); i++ {
			if wMap[ns[:i]] {
				DFS(ns[i:], append(r, ns[:i]))
			}
		}
	}
	DFS(s, []string{})
	return re
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	//s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	//wordDict = []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	s = "pineapplepenapple"
	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	res := wordBreak8(s, wordDict)
	for _, r := range res {
		fmt.Println(r)
	}
}

/* vim: set tabstop=4 set shiftwidth=4 */
