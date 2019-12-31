package main

import "fmt"

func main() {

}

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{
			nums,
		}
	}
	res := [][]int{}

	var helper func(nums []int, left, right int)
	helper = func(nums []int, left, right int) {
		if left == right {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := left; i <= right; i++ {
			nums[left], nums[i] = nums[i], nums[left]
			helper(nums, left+1, right)
			nums[left], nums[i] = nums[i], nums[left]
		}
	}

	helper(nums, 0, len(nums)-1)
	fmt.Println(res)
	return res
}
