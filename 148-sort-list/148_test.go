package main

import "testing"

func TestOk(t *testing.T) {
	arr := []int{
		9, 13, 5, 6, 8, 3, 5, 2, 11,
	}
	head := build(arr)
	showListNode(head)
	sortList(head)
	showListNode(head)
}
