#!/usr/bin/env python
# We consider array1[i] to be the first element in array1
# We consider array2[x] to be the first element in array2
def getKth(array1, i, array2, x, k):
    print(array1,i,array2,x,k)
    if i == len(array1):
        return array2[x + k]
    elif x == len(array2):
        return array1[i + k]
    elif k == 0:
        return min(array1[i], array2[x])

    mid1 = min(len(array1) - i, (k + 1) // 2)
    mid2 = min(len(array2) - x, (k + 1) // 2)
    a = array1[i + mid1 - 1]
    b = array2[x + mid2 - 1]
    print(mid1,mid2,a,b)

    if a < b:
        return getKth(array1, i + mid1, array2, x, k - mid1)
    return getKth(array1, i, array2, x + mid2, k - mid2)

# This function assumes that we have at least 1 number in the array.
def findMedianSortedArrays(nums1, nums2):
    total_nums = len(nums1) + len(nums2)
    midpoint = total_nums // 2 + 1

    if total_nums % 2 == 0:
        first = getKth(nums1, 0, nums2, 0, total_nums // 2 - 1)
        print("---")
        second = getKth(nums1, 0, nums2, 0, total_nums // 2)
        return (first + second) / 2
    else:
        return getKth(nums1, 0, nums2, 0, total_nums // 2)  

nums1 = [1,2]
nums2 = [-1,3]

print(findMedianSortedArrays(nums1,nums2))
