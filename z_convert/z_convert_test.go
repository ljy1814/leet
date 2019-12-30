package main

import "testing"

func Test1(t *testing.T) {
	src := "LEETCODEISHIRING"
	numRows := 3
	want := "LCIRETOESIIGEDHN"

	got := convert(src, numRows)

	if got != want {
		t.Fatalf("convert got=%q,want=%q", got, want)
	}
}

func Test2(t *testing.T) {
	src := "LEETCODEISHIRING"
	numRows := 4
	want := "LDREOEIIECIHNTSG"

	got := convert(src, numRows)

	if got != want {
		t.Fatalf("convert got=%q,want=%q", got, want)
	}
}
