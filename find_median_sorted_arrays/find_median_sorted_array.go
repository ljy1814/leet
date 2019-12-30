package main

import (
	"fmt"
	"math"
	"mlog"
)

const maxUint = ^uint(0)
const minUint = 0
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

/*
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0
*/

/*
nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	if nums1 == nil || len(nums1) <= 0 {
		if nums2 != nil && len(nums2) > 0 {
			if len(nums2)%2 == 0 {
				return float64(nums2[len(nums2)>>1]+nums2[len(nums2)>>1-1]) / 2
			}
			return float64(nums2[len(nums2)>>1])
		}
		return 0.0
	}

	if nums2 == nil || len(nums2) <= 0 {
		if nums1 != nil && len(nums1) > 0 {
			if len(nums1)%2 == 0 {
				return float64(nums1[len(nums1)>>1]+nums1[len(nums1)>>1-1]) / 2
			}
			return float64(nums1[len(nums1)>>1])
		}
	}

	// nums1长
	if len(nums1) < len(nums2) {
		nums2, nums1 = nums1, nums2
	}

	l1 := len(nums1) - 1
	l2 := len(nums2) - 1
	l3 := l1 + l2
	half := (l3 + 1) >> 1
	for i := half; i <= l1; {
		j := half - i

		//fmt.Println(i, nums1[i], j, nums2[j], l3, half)
		fmt.Println(i, j, l3, half)

		if i > 0 && j <= l2-1 && nums1[i] > nums2[j+1] {
			//fmt.Println("++++++", i, j)
			i--
		} else if i <= l1-1 && j >= 0 && nums1[i+1] < nums2[j] {
			//fmt.Println("~~~~~", i, j)
			i++
		} else {
			// 左半部最大,又半部分最小
			fmt.Println("++++++", i, j)
			//fmt.Println("----", i, nums1[i], j, nums2[j], l3, half)
			for j < 0 {
				j++
				i--
			}

			if l3%2 == 1 {
				if nums1[i] > nums2[j] {
					return float64(max(nums1[i-1], nums2[j]))
				}

				if j == 0 {
					if nums1[i] < nums2[j] {
						return float64(nums1[i])
					}
					return float64(max(nums1[i-1], nums2[j]))
				}

				return float64(max(nums2[j-1], nums1[i]))
			}

			if j == 0 && i == 0 {
				return float64(nums1[i]+nums2[j]) / 2
			}

			if j == 0 {
				if nums1[i] > nums2[j] {
					return float64(nums1[i]+max(nums2[j], nums1[i-1])) / 2
				}

				if i == l1 {
					return float64(nums1[i]+nums2[j]) / 2
				}

				return float64(nums1[i]+min(nums1[i+1], nums2[j])) / 2
			}

			if i == 0 {
				if nums1[i] > nums2[j] {
					return float64(nums1[i]+nums2[j]) / 2
				}

				return float64(nums2[j]+max(nums2[j-1], nums1[i])) / 2
			}

			if nums1[i] > nums2[j] {
				return float64(nums1[i]+max(nums2[j], nums1[i-1])) / 2
			}

			return float64(nums2[j]+max(nums2[j-1], nums1[i])) / 2

			/*
				if j == 0 && i < l1-1 && nums1[i+1] < nums2[j] {
					if l3%2 == 1 {
						return float64(nums1[i])
					}
					return float64(nums1[i]+nums1[i+1]) / 2
				} else if j == l2 && i > 0 && nums1[i-1] > nums2[j] {
					if l3%2 == 1 {
						return float64(nums1[half-j-1])
					}
					return float64(nums1[half-j]+nums1[half-j-1]) / 2.0
				} else {
					fmt.Println("----", i, j, l3, half)
					if l3%2 == 1 {
						if nums1[i] > nums2[j] {
							return float64(nums2[j])
						}
					}
					return float64(nums1[i]+nums2[j]) / 2.0
				}
			*/
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getKth(nums1 []int, i int, nums2 []int, j, k int) int {
	fmt.Println(i, j, k)
	if i == len(nums1) && j+k < len(nums2) {
		return nums2[j+k]
	} else if j == len(nums2) && i+k < len(nums1) {
		return nums1[i+k]
	} else if k == 0 && i < len(nums1) && j < len(nums2) {
		return min(nums1[i], nums2[j])
	}

	mid1 := min(len(nums1)-i, (k+1)/2)
	mid2 := min(len(nums2)-j, (k+1)/2)

	a := nums1[i+mid1-1]
	b := nums2[j+mid2-1]

	fmt.Println(mid1, mid2, a, b)

	if a < b {
		return getKth(nums1, i+mid1, nums2, j, k-mid1)
	}

	return getKth(nums1, i, nums2, j+mid2, k-mid2)
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	totalNums := len(nums1) + len(nums2)
	//midPoint := totalNums>>1 + 1

	if totalNums%2 == 0 {
		first := getKth(nums1, 0, nums2, 0, totalNums>>1-1)
		second := getKth(nums1, 0, nums2, 0, totalNums>>1)

		return float64(first+second) / 2
	}

	return float64(getKth(nums1, 0, nums2, 0, totalNums>>1))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n1 := len(nums1)
	n2 := len(nums2)
	fmt.Println(mlog.GetLineNumber(), n1, n2)
	if len(nums2) == 0 {
		if n1%2 == 0 {
			return float64(nums1[n1>>1]+nums1[n1>>1+1]) / 2
		}
		return float64(nums1[n1>>1])
	}

	// 虚拟数组映射,补空
	/*
		A,B,C -> *,A,*,B,*,C,*
		n->2*n+1
	*/
	lo := 0
	hi := 2 * n2

	var (
		l2, r2 int
		l1, r1 int
		c1     int
		c2     int
	)

	for lo <= hi {
		c2 = (hi-lo)>>1 + lo
		c1 = n1 + n2 - c2
		fmt.Println(mlog.GetLineNumber(), lo, hi, c1, c2)

		if c2>>1 >= n2 {
			r2 = maxInt
		} else {
			r2 = nums2[c2>>1]
		}

		if c2 <= 0 {
			l2 = int(math.MinInt64)
		} else {
			l2 = nums2[(c2-1)>>1]
		}

		l1 = nums1[(c1-1)>>1]
		if c1>>1 >= n1 {
			r1 = maxInt
		} else {
			r1 = nums1[c1>>1]
		}

		fmt.Println(mlog.GetLineNumber(), l1, r1, l2, r2)
		if l1 > r2 {
			lo = c2 + 1
			continue
		}
		if l2 > r1 {
			hi = c2 - 1
			continue
		}
		break
	}

	ml := l1
	if l2 > ml {
		ml = l2
	}

	mr := r1
	if r2 < mr {
		mr = r2
	}

	fmt.Println(mlog.GetLineNumber(), ml, mr)

	return float64(ml+mr) / 2
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3, 5, 7, 9, 11}
	nums2 = []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3, 5, 7, 9, 11}
	nums2 = []int{2, 4, 6, 10}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3, 5, 7, 9, 11}
	nums2 = []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums2 = []int{1, 3, 5, 7, 9, 11}
	nums1 = []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums2 = []int{1, 2, 3, 4, 6}
	nums1 = []int{5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
