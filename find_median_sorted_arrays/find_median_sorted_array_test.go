package main

import (
	"testing"
)

func Test1(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-1, 3}
	want := 1.5
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test2(t *testing.T) {

	nums2 := []int{1, 2, 3, 4, 5, 7, 8}
	nums1 := []int{6}
	want := 4.5
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test3(t *testing.T) {

	nums2 := []int{1, 2, 3, 4, 5, 7}
	nums1 := []int{6}
	want := 4.0
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test4(t *testing.T) {

	want := 4.5
	nums2 := []int{1, 2, 3, 4, 5, 8}
	nums1 := []int{6, 7}
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test5(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	want := 2.5
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test6(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{1, 1}
	want := 1.0
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test7(t *testing.T) {
	nums1 := []int{2, 3}
	nums2 := []int{1}
	want := 2.0
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test8(t *testing.T) {
	nums1 := []int{1, 4}
	nums2 := []int{2, 3}
	want := 2.5
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}

func Test9(t *testing.T) {
	nums1 := []int{-1, 1, 4, 6, 8, 9, 12, 17}
	nums2 := []int{-1, 2, 3, 5, 7, 13}
	want := 5.5
	got := findMedianSortedArrays(nums1, nums2)
	if got != want {
		t.Fatalf("findMedianSortedArrays got=%f want=%f", got, want)
	}
}
