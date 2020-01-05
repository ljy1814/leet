package main

import "fmt"

func main() {
	s := ")())("
	s = "()())()"
	fmt.Println(removeInvalidParentheses(s))
}

func removeInvalidParentheses(s string) []string {
	check := func(s string) bool {
		cnt := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				cnt++
			} else if s[i] == ')' {
				cnt--
				if cnt < 0 {
					return false
				}
			}
		}
		return cnt == 0
	}

	res := []string{}

	queue := make([]string, 0)
	visited := make(map[string]bool)
	queue = append(queue, s)
	maxn := 0
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			str := queue[i]
			if visited[str] {
				continue
			}
			visited[str] = true

			if check(str) {
				if len(str) > maxn {
					maxn = len(str)
				}
				res = append(res, str)
				continue
			}

			for j := 0; j < len(str); j++ {
				if s[j] != ')' && s[j] != '(' {
					continue
				}
				queue = append(queue, str[:j]+str[j+1:])
			}
		}
		tmp := make([]string, len(queue)-n)
		copy(tmp, queue[n:])
		queue = nil
		queue = tmp
	}

	ret := []string{}
	for _, str := range res {
		if len(str) == maxn {
			ret = append(ret, str)
		}
	}

	return ret
}

func removeInvalidParentheses1(s string) []string {
	check := func(s string) bool {
		cnt := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				cnt++
			} else if s[i] == ')' {
				cnt--
				if cnt < 0 {
					return false
				}
			}
		}
		return cnt == 0
	}

	res := []string{}

	var dfs func(s string, st, l, r int)
	dfs = func(s string, st, l, r int) {
		if l == 0 && r == 0 {
			if check(s) {
				res = append(res, s)
			}
			return
		}

		for i := st; i < len(s); i++ {
			// 重复 (( 删除其中任意一个得到的结果是相同的
			if i-1 >= st && s[i] == s[i-1] {
				continue
			}
			if l > 0 && s[i] == '(' {
				dfs(s[:i]+s[i+1:], i, l-1, r)
			}
			if r > 0 && s[i] == ')' {
				dfs(s[:i]+s[i+1:], i, l, r-1)
			}
		}
	}

	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	dfs(s, 0, left, right)
	return res
}
