package main

import (
	"container/heap"
	"sort"
)

func main() {
}

type LessQueue []int

// 堆顺序和>相同
func (pq LessQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq LessQueue) Len() int {
	return len(pq)
}

func (pq LessQueue) Head() int {
	if len(pq) <= 0 {
		return 0
	}
	return pq[0]
}

func (pq LessQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *LessQueue) Clear() {
	*pq = []int{}
}

func (pq *LessQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *LessQueue) Pop() interface{} {
	n := len(*pq)
	t := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return t
}

type MoreQueue []int

func (pq MoreQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq MoreQueue) Len() int {
	return len(pq)
}

func (pq MoreQueue) Head() int {
	if len(pq) <= 0 {
		return 0
	}
	return pq[0]
}

func (pq MoreQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MoreQueue) Clear() {
	*pq = []int{}
}

func (pq *MoreQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MoreQueue) Pop() interface{} {
	n := len(*pq)
	t := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return t
}

// 超时  效果不如sort
func medianSlidingWindow1(nums []int, k int) []float64 {
	n := len(nums)
	if k > n {
		return nil
	}

	ret := []float64{}
	for i := 0; i+k-1 < n; i++ {
		lq := &LessQueue{}
		mq := &MoreQueue{}

		for j := 0; j < k; j++ {
			nn := nums[i+j]
			if lq.Len() == 0 || lq.Head() > nn {
				heap.Push(lq, nn)
			} else {
				heap.Push(mq, nn)
			}

			if mq.Len() > lq.Len() {
				heap.Push(lq, heap.Pop(mq))
			} else if lq.Len() > mq.Len()+1 {
				heap.Push(mq, heap.Pop(lq))
			}
		}

		m := 0.0
		if lq.Len() > mq.Len() {
			m = float64(lq.Head())
		} else {
			m = float64(lq.Head()+mq.Head()) / 2
		}
		lq.Clear()
		mq.Clear()
		ret = append(ret, m)
	}

	return ret
}

func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	if k > n {
		return nil
	}

	getMid := func(nums []int) float64 {
		n := len(nums)
		t := nums[n>>1] + nums[(n-1)>>1]
		return float64(t) / 2
	}

	binarySearch := func(nums []int, k int) int {
		l, r := 0, len(nums)-1
		for l < r {
			m := l + (r-l)>>1
			if nums[m] > k {
				r = m - 1
			} else if nums[m] < k {
				l = m + 1
			} else {
				return m
			}
		}
		return l
	}

	ret := make([]float64, n-k+1)
	knums := make([]int, k)
	copy(knums, nums)
	sort.Ints(knums)
	ret[0] = getMid(knums)

	for i := k; i < n; i++ {
		index := binarySearch(knums, nums[i-k])
		tmp := append([]int{}, knums[:index]...)
		tmp = append(tmp, knums[index+1:]...)
		index = binarySearch(tmp, nums[i])

		if index < k-1 && tmp[index] < nums[i] {
			knums = append([]int{}, tmp[:index+1]...)
			knums = append(knums, nums[i])
			knums = append(knums, tmp[index+1:]...)
		} else {
			knums = append([]int{}, tmp[:index]...)
			knums = append(knums, nums[i])
			knums = append(knums, tmp[index:]...)
		}

		tmp = nil

		ret[i-k+1] = getMid(knums)
	}
	return ret
}
