package main

import (
	"container/heap"
	"testing"
)

func buildLists() []*ListNode {
	lns := []*ListNode{}
	arr := []int{
		1, 4, 5,
	}
	ln := build(arr)
	lns = append(lns, ln)
	arr = []int{
		1, 3, 4,
	}
	ln = build(arr)
	lns = append(lns, ln)

	arr = []int{
		2, 6,
	}
	ln = build(arr)
	lns = append(lns, ln)

	arr = []int{
		3, 7, 9, 13, 16,
	}
	ln = build(arr)
	lns = append(lns, ln)

	return lns
}

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lns := buildLists()
		pq := &PriorityQueue{}
		for _, l := range lns {
			if l != nil {
				heap.Push(pq, l)
			}
		}
		dummy := &ListNode{}
		cur := dummy

		for pq.Len() > 0 {
			t := heap.Pop(pq)
			x := t.(*ListNode)
			if x.Next != nil {
				heap.Push(pq, x.Next)
			}
			cur.Next = x
			cur = cur.Next
		}
	}
}

func Benchmark2Merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lns := buildLists()
		mergeKLists(lns)
	}
}
