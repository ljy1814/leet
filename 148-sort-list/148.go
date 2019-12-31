package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func build(arr []int) *ListNode {
	if len(arr) <= 0 {
		return nil
	}
	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}
	cur := head
	for i := 1; i < len(arr); i++ {
		t := &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		cur.Next = t
		cur = cur.Next
	}
	return head
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sortList1(head, nil)
	return head
}
func sortList1(head, tail *ListNode) {
	if head == tail {
		return
	}
	v := head.Val
	cur := head.Next
	p := head
	for cur == tail && cur != nil || cur != nil {
		if cur.Val < v {
			p = p.Next
			p.Val, cur.Val = cur.Val, p.Val
		}
		cur = cur.Next
	}

	p.Val, head.Val = head.Val, p.Val
	sortList1(head, p)
	sortList1(p.Next, tail)
}

func showListNode(ln *ListNode) {
	if ln == nil {
		return
	}

	for ln != nil {
		fmt.Printf("%d ", ln.Val)
		ln = ln.Next
	}
	fmt.Println()
}

func main() {
}
