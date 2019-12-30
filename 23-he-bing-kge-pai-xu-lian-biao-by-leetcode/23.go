package main

import (
	"container/heap"
	"fmt"
)

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	if n <= 0 {
		return nil
	}
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}

	for len(lists) > 1 {
		tmpList := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			if i < len(lists)-1 {
				tmpList = append(tmpList, merge2Lists(lists[i], lists[i+1]))
			} else {
				tmpList = append(tmpList, merge2Lists(lists[i], nil))
			}
		}
		lists = tmpList
	}
	return lists[0]
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		cur = l1
		l1 = l1.Next
	} else {
		cur = l2
		l2 = l2.Next
	}
	head = cur

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head
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
	lns := []*ListNode{}
	arr := []int{
		1, 4, 5,
	}
	ln := build(arr)
	showListNode(ln)
	lns = append(lns, ln)
	arr = []int{
		1, 3, 4,
	}
	ln = build(arr)
	lns = append(lns, ln)
	showListNode(ln)
	fmt.Println()
	//showListNode(merge2Lists(lns[0], lns[1]))

	fmt.Println("------")
	arr = []int{
		2, 6,
	}
	ln = build(arr)
	showListNode(ln)
	lns = append(lns, ln)

	arr = []int{
		3, 7, 9, 13, 16,
	}
	ln = build(arr)
	showListNode(ln)
	lns = append(lns, ln)
	//showListNode(mergeKLists(lns))

	// priority queue
	pq := &PriorityQueue{}
	for _, l := range lns {
		if l != nil {
			heap.Push(pq, l)
		}
	}
	//heap.Init(pq)
	dummy := &ListNode{}
	cur := dummy

	fmt.Println("------")
	for pq.Len() > 0 {
		t := heap.Pop(pq)
		fmt.Println("Pop t:", t)
		x := t.(*ListNode)
		if x.Next != nil {
			heap.Push(pq, x.Next)
		}
		cur.Next = x
		cur = cur.Next
	}
	showListNode(dummy.Next)
}
