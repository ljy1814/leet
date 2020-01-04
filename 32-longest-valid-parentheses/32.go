package main

import "fmt"

type Stack struct {
	stack   []int
	top     int
	capcity int
}

func (s *Stack) push(i int) {
	if s.top == s.capcity-1 {
		s.capcity++
		s.top++
		s.stack = append(s.stack, i)
		return
	}
	s.top++
	s.stack[s.top] = i
}

func (s *Stack) pop() int {
	t := s.stack[s.top]
	s.top--

	if s.capcity>>1 > s.top && s.top > 1 {
		st := make([]int, s.top+1)
		copy(st, s.stack)
		s.stack = st
		s.capcity = len(st)
	}

	return t
}

func (s *Stack) peak() int {
	return s.stack[s.top]
}

func (s *Stack) empty() bool {
	return s.top < 0
}

func longestValidParentheses1(s string) int {
	if len(s) <= 1 {
		return 0
	}
	n := len(s)
	m := 0
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	st := &Stack{top: -1}
	// 上一次匹配失败的位置
	st.push(-1)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			st.push(i)
		} else {
			st.pop()

			// 出现有问题的)
			// 连续))) 以最后一个为最后一次匹配失败的位置
			// )()())
			if st.empty() {
				st.push(i)
			} else {
				// (((()))
				// ((((i
				// ((( i
				m = max(m, i-st.peak())
			}
		}
		fmt.Println(st.stack)
	}

	return m
}

func main() {
}

func longestValidParentheses(s string) int {
	m := 0
	// 左右括号数量
	left, right := 0, 0
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		// 剔除毛刺 )))()  不匹配
		if right > left {
			right = 0
			left = 0
		} else if right == left {
			m = max(m, 2*right)
		} else {
			// 其它 ((()) // 左边大于右边,只能通过从右往左再扫描一遍得到正确结果
			// ((()) 得到的结果为0
		}
	}

	left, right = 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		// 剔除毛刺((((())不匹配
		if left > right {
			right = 0
			left = 0
		} else if right == left {
			m = max(m, 2*right)
		}
		// )))(())()(((()
	}
	return m
}
