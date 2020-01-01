package main

func main() {
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	dm := make(map[int]int)
	w := t + 1
	n := len(nums)

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	getId := func(x, w int) int {
		if x < 0 {
			return x / w
		}
		return x/w + 1
	}

	for i := 0; i < n; i++ {
		m := getId(nums[i], w)
		if _, ok := dm[m]; ok {
			return true
		}
		if _, ok := dm[m-1]; ok && abs(nums[i]-dm[m-1]) < w {
			return true
		}
		if _, ok := dm[m+1]; ok && abs(nums[i]-dm[m+1]) < w {
			return true
		}

		dm[m] = nums[i]
		if i >= k {
			delete(dm, getId(nums[i-k], w))
		}
	}

	return false
}
