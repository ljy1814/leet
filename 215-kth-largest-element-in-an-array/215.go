package main

import (
	"container/heap"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	part := func(nums []int, left, right int) int {
		v := nums[left]
		p := left

		for j := left; j < right; j++ {
			if nums[j] > v {
				p++
				nums[p], nums[j] = nums[j], nums[p]
			}
		}
		nums[p], nums[left] = nums[left], nums[p]
		return p
	}

	left := 0
	right := len(nums)
	for left < right {
		p := part(nums, left, right)
		if p == k-1 {
			return nums[p]
		} else if p < k-1 {
			left = p + 1
		} else {
			right = p
		}
	}
	return nums[k-1]
}

type (
	PriorityQueue []int
)

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[j] < pq[i]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	t := old[n-1]
	old = nil

	*pq = (*pq)[:n-1]
	return t
}

func findKthLargest1(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	pq := &PriorityQueue{}

	for _, n := range nums {
		heap.Push(pq, n)
	}

	var x interface{}
	for i := 0; i < k && pq.Len() > 0; i++ {
		x = heap.Pop(pq)
		fmt.Println(i, x)
	}

	return x.(int)
}

func main() {
}
