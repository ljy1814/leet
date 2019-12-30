package main

import "fmt"

/*
LEETCODEISHIRING

L   C   I   R
E T O E S I I G
E   D   H   N

LCIRETOESIIGEDHN

L     D     R
E   O E   I I
E C   I H   N
T     S     G
LDREOEIIECIHNTSG
*/

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	cell := 2*numRows - 2
	lists := make([][]byte, numRows)

	for i := 0; i < len(s); i++ {
		m := i % cell
		index := m
		if index > numRows-1 {
			index = cell - index
		}
		if lists[index] == nil {
			lists[index] = make([]byte, 0)
		}

		lists[index] = append(lists[index], s[i])
	}

	ret := ""
	for i := 0; i < len(lists); i++ {
		ret += string(lists[i])
	}

	return ret
}

func main() {
	fmt.Println("vim-go")
}
