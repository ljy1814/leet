package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers0(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return &ListNode{}
	}
	carry := 0

	p1 := l1
	p2 := l2
	ret := &ListNode{}
	p3 := ret
	cur := p3
	for p1 != nil && p2 != nil {
		if p3 == nil {
			p3 = &ListNode{}
		}

		p3.Val = p1.Val + p2.Val + carry
		fmt.Printf("p1=%d,p2=%d,p3=%d,carry=%d\n", p1.Val, p2.Val, p3.Val, carry)
		if p3.Val >= 10 {
			carry = 1
			p3.Val = p3.Val % 10
		} else {
			carry = 0
		}
		fmt.Printf("----p1=%d,p2=%d,p3=%d,carry=%d\n", p1.Val, p2.Val, p3.Val, carry)

		p1 = p1.Next
		p2 = p2.Next
		cur = p3
		p3.Next = &ListNode{}
		p3 = p3.Next
	}

	for p1 != nil {
		p3.Val = p1.Val + carry
		if p3.Val >= 10 {
			carry = 1
			p3.Val = p3.Val % 10
		} else {
			carry = 0
		}

		p1 = p1.Next
		cur = p3
		p3.Next = &ListNode{}
		p3 = p3.Next
	}

	for p2 != nil {
		p3.Val = p2.Val + carry
		if p3.Val >= 10 {
			carry = 1
			p3.Val = p3.Val % 10
		} else {
			carry = 0
		}
		p2 = p2.Next
		cur = p3
		p3.Next = &ListNode{}
		p3 = p3.Next
	}

	if carry > 0 {
		p3.Val = carry
	}

	if p3.Val == 0 {
		cur.Next = nil
	}

	return ret
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	p1 := l1
	p2 := l2
	ret := &ListNode{}
	p3 := ret
	cur := p3
	for p1 != nil || p2 != nil || carry > 0 {
		s := 0
		if p1 != nil {
			s += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			s += p2.Val
			p2 = p2.Next
		}

		s += carry
		carry = s / 10
		s = s % 10

		p3.Val = s
		p3.Next = &ListNode{}
		cur = p3
		p3 = p3.Next
	}
	cur.Next = nil

	return ret
}

func getResult(l *ListNode) int {
	ret := 0
	base := 1
	for l != nil {
		fmt.Println(l.Val)
		ret = ret + l.Val*base
		l = l.Next
		base *= 10
	}

	return ret
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l3 := addTwoNumbers(l1, l2)
	fmt.Println(getResult(l3))

	l1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	l2 = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	fmt.Println(getResult((addTwoNumbers(l1, l2))))

	l2 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	l1 = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	fmt.Println(getResult((addTwoNumbers(l1, l2))))

	l1 = &ListNode{
		Val:  5,
		Next: nil,
	}
	l2 = &ListNode{
		Val:  5,
		Next: nil,
	}
	fmt.Println(getResult((addTwoNumbers(l1, l2))))

	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  8,
			Next: nil,
		},
	}
	l2 = &ListNode{
		Val:  0,
		Next: nil,
	}
	fmt.Println(getResult((addTwoNumbers(l1, l2))))
}
