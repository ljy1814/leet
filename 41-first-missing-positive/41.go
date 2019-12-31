package main

// 桶排序

func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}

	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && i != nums[i]-1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
		/*
			for nums[i] != i+1 {
				idx := nums[i] - 1
				// [idx] == [i] -> 重复
				if idx < 0 || idx >= n || nums[i] == nums[idx] {
					break
				}
				nums[i], nums[idx] = nums[idx], nums[i]
			}
		*/
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
}
