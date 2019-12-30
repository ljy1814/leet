package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) <= 0 {
		return []int{0, 0}
	}
	nm := map[int]int{}
	used := map[int]bool{}

	for i, n := range nums {
		nm[n] = i
		used[i] = false
	}

	for i, n := range nums {
		used[i] = true
		if m, ok := nm[target-n]; ok && !used[m] {
			if target-n == n {
				for j, t := range nums {
					if n == t && j != i {
						return []int{
							i, j,
						}
					}
				}
			}
			return []int{
				i, nm[target-n],
			}
		}
		used[i] = false
	}

	return []int{0, 0}
}
func twoSum2(nums []int, target int) []int {
	vm := map[int]int{}

	for k, v := range nums {
		if k1, ok := vm[target-v]; ok {
			return []int{k, k1}
		}
		vm[v] = k
	}
}

func main() {
	fmt.Println("vim-go")
	nums := []int{
		2, 7, 11, 15,
	}

	target := 9
	fmt.Println(twoSum(nums, target))
	target = 18
	fmt.Println(twoSum(nums, target))
	nums = []int{
		3, 3,
	}
	target = 6
	fmt.Println(twoSum(nums, target))

	nums = []int{
		3, 2, 4,
	}
	target = 6
	fmt.Println(twoSum(nums, target))
}
