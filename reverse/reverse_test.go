package main

import (
	"testing"
)

/*
 * Author : yajin
 * Email : yajin160305@gmail.com
 * File : reverse_test.go
 * CreateDate : 2018-12-22 23:23:28
 * */

func Test1(t *testing.T) {
	x := 123
	want := 321
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}

func Test2(t *testing.T) {
	x := 120
	want := 21
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}

func Test3(t *testing.T) {
	x := -120
	want := -21
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}

func Test4(t *testing.T) {
	x := -4294967296
	want := 0
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}

func Test5(t *testing.T) {
	x := 4294967293
	want := 3927694924
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}

func Test6(t *testing.T) {
	x := 4294967206
	want := 0
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}

func Test7(t *testing.T) {
	x := 32768
	want := 86723
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}

func Test8(t *testing.T) {
	x := 1534236469
	want := 0
	got := reverse(x)
	if got != want {
		t.Fatalf("got=%d, want=%d", got, want)
	}
}
/* vim: set tabstop=4 set shiftwidth=4 */
