package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int, sum int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	fmt.Println(nums)
	ret := make(map[string][]int)
	dup := make(map[string]bool)

	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1

		for l < r {
			s := nums[l] + nums[r] + nums[i]
			if s == sum {
				key := fmt.Sprintf("%d:%d:%d", i, l, r)
				dk := fmt.Sprintf("%d:%d:%d", nums[i], nums[l], nums[r])
				if ret[key] == nil && !dup[dk] {
					ret[key] = []int{
						nums[i], nums[l], nums[r],
					}
					dup[dk] = true
				}
				l0 := l
				r0 := r
				for r > l && nums[r0] == nums[r-1] {
					r--
				}
				r--
				for l < r && nums[l0] == nums[l+1] {
					l++
				}
				l++
			} else if s < sum {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			} else {
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				r--
			}
		}
	}

	rets := [][]int{}
	for k, r := range ret {
		fmt.Println(k, r)
		rets = append(rets, r)
	}

	return rets
}

func threeSum1(nums []int, sum int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	fmt.Println(nums)
	rets := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1

		for l < r {
			s := nums[l] + nums[r] + nums[i]
			if s == sum {
				rets = append(rets, []int{
					nums[i], nums[l], nums[r],
				})
				fmt.Println(i, l, r)
				l0 := l
				r0 := r
				for r > l && nums[r0] == nums[r-1] {
					r--
				}
				r--
				for l < r && nums[l0] == nums[l+1] {
					l++
				}
				l++
			} else if s < sum {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			} else {
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				r--
			}
		}

		for ; i < len(nums)-3 && nums[i] == nums[i+1]; i++ {

		}
	}

	return rets
}

func main() {
	nums := []int{
		-1, 0, 1, 2, -1, -4,
	}
	sum := 0

	fmt.Println(threeSum1(nums, sum))
}
