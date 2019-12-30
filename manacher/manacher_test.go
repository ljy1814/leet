package main

import (
	"bytes"
	"fmt"
	"mlog"
	"testing"
)

const (
	// MAX_SIZE max
	MaxSize = 1024
)

func mancher1(str string) string {
	if str == "" || len(str) == 1 {
		return str
	}

	center := 0
	rm := 0
	//	pos := make([]int, len(str), len(str))
	pos := [MaxSize]int{}
	pos[0] = 0

	for i := 1; i < len(str); i++ {
		pos[i] = 1
		rm = center + pos[center]
		iMirror := center - (i - center)

		if i > center && iMirror >= 0 {
			if rm > i+pos[iMirror] {
				pos[i] = pos[iMirror]
			} else {
				pos[i] = rm - i
			}
		}

		for {
			if i+pos[i] < len(str) && i-pos[i] >= 0 {
				if pos[i] == 0 {
					pos[i]++
					continue
				}
				fmt.Printf("%d+pos[%d]=%d:%c,%d-pos[%d]=%d:%c\n", i, i, i+pos[i], str[i+pos[i]], i, i, i-pos[i], str[i-pos[i]])

				if str[i+pos[i]] == str[i-pos[i]] {
					//	i+1+pos[i] < len(str) && str[i+1+pos[i]] == str[i-pos[i]]
					pos[i]++
					continue
				}
			}
			break
		}

		if i+pos[i] > rm {
			rm = i + pos[i]
			center = i
		}
		if rm >= len(str)-1 {
			break
		}
	}

	maxPos := 0
	maxCount := 0
	for i := 0; i < len(pos); i++ {
		if pos[i] > maxCount {
			maxCount = pos[i]
			maxPos = i
		}
	}

	fmt.Println(maxPos, maxCount)
	return str[maxPos-maxCount+1 : maxPos+maxCount]
}

func extendManacher(str string) string {
	if str == "" {
		return str
	}

	ret := bytes.Buffer{}
	ret.WriteByte('^')
	for i := 0; i < len(str); i++ {
		ret.WriteByte('#')
		ret.WriteByte(str[i])
		if i == len(str)-1 {
			ret.WriteByte('#')
			ret.WriteByte('$')
		}
	}

	return ret.String()
}

func mancher2(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}

	str := extendManacher(s)

	center := 0
	rm := 0
	pos := make([]int, len(str), len(str))
	pos[0] = 0

	for i := 1; i < len(str); i++ {
		pos[i] = 1
		rm = center + pos[center]
		iMirror := center - (i - center)

		if i > center && iMirror >= 0 {
			if rm > i+pos[iMirror] {
				pos[i] = pos[iMirror]
			} else {
				pos[i] = rm - i
			}
		}

		for {
			if i+pos[i] < len(str) && i-pos[i] >= 0 {
				if pos[i] == 0 {
					pos[i]++
					continue
				}
				fmt.Printf("%d+pos[%d]=%d:%c,%d-pos[%d]=%d:%c\n", i, i, i+pos[i], str[i+pos[i]], i, i, i-pos[i], str[i-pos[i]])

				if str[i+pos[i]] == str[i-pos[i]] {
					pos[i]++
					continue
				}
			}
			break
		}

		if i+pos[i] > rm {
			rm = i + pos[i]
			center = i
		}
		if rm >= len(str)-1 {
			break
		}
	}

	maxPos := 0
	maxCount := 0
	for i := 0; i < len(pos); i++ {
		if pos[i] > maxCount {
			maxCount = pos[i]
			maxPos = i
		}
	}

	fmt.Println(str, maxPos, maxCount)
	return s[(maxPos-maxCount+1)>>1 : (maxPos+maxCount-1)>>1]
}

func PlainMatch(str string) string {
	if str == "" || len(str) == 1 {
		return str
	}

	l, r := 0, 0
	for i := 0; i < len(str); i++ {
		for j := 1; j+i < len(str); j++ {
			if i-j > 0 && i+j < len(str) && str[i+j] == str[i-j] {
				continue
			}
			if i+j-1-(i-j+1) > r-l {
				l = i - j + 1
				r = i + j - 1
			}
			break
		}

		for j := 1; j+i < len(str); j++ {
			if i-j > 0 && i+1+j < len(str) && str[i+1+j] == str[i-j] {
				continue
			}
			if i+1+j-1-(i-j+1) > r-l {
				l = i - j + 1
				r = i + 1 + j - 1
			}
			break
		}
	}

	return str[l : r+1]
}

func TestManacher1(t *testing.T) {
	str := "abaxabaxabybaxabyb"
	fmt.Println(len(str))
	fmt.Println(mancher1(str))
}

func TestPlain(t *testing.T) {
	str := "abaxabaxabybaxabyb"
	str = "abaxabaxabyybaxabyb"
	fmt.Println(len(str))
	fmt.Println(PlainMatch(str))
}

func BenchmarkManacher1(b *testing.B) {
	str := "abaxabaxabybaxabyb"
	for i := 0; i < b.N; i++ {
		mancher1(str)
	}
}

func BenchmarkPlain(b *testing.B) {
	str := "abaxabaxabybaxabyb"
	str = "abaxabaxabyybaxabyb"
	for i := 0; i < b.N; i++ {
		PlainMatch(str)
	}
}

func TestManacher2(t *testing.T) {
	str := "abaxabaxabybaxabyb"
	fmt.Println(mancher2(str))
	fmt.Println(mancher2(str))
	str = "abaxabaxabyybaxabyb"
	fmt.Println(mancher1(str))
	fmt.Println(mancher2(str))

	fmt.Println("---------")
	str = "babad"
	fmt.Println(mancher1(str))
	fmt.Println(mancher2(str))

	fmt.Println("---------")
	str = "cbbd"
	fmt.Println(mancher1(str))
	fmt.Println(mancher2(str))

	fmt.Println("---------")
	str = "ac"
	fmt.Println(mancher2(str))

	fmt.Println("---------")
	str = "bb"
	fmt.Println(mancher2(str))
}

func mancher3(str string) string {
	if len(str) <= 1 {
		return str
	}
	var (
		// C center
		C int
		// R right
		R       int
		m       int
		iMirror int
		maxI    int
		maxM    int
		n       = len(str)
		nn      = 2 * len(str)
		l       int
		r       int
	)
	P := make([]int, nn+1)

	for i := 1; i < nn; i++ {
		iMirror = 2*C - i
		m = 1

		if iMirror >= 0 && P[iMirror] < m {
			m = P[iMirror]
		}

		fmt.Println(mlog.GetLineNumber(), i, iMirror, m, l, r, n)
		//fmt.Println(mlog.GetLineNumber(), i, m, l, r, n, str[l], str[r])
		for i-m >= 0 && i+m <= nn {
			if (i+m)%2 == 0 {
				m++
				continue
			}
			l = (i - m) >> 1
			r = (i + m - 1) >> 1
			fmt.Println(mlog.GetLineNumber(), l, r, i, m)
			if str[l] == str[r] {
				m++
				continue
			}
			break
		}

		P[i] = m

		if m > maxM {
			maxM = m
			maxI = i
		}

		if i+m > R {
			C = i
			R = i + m
		}
		fmt.Println(mlog.GetLineNumber(), i, P[i], C, R, m, maxI, maxM)

		if R >= nn {
			break
		}
	}

	l = maxI - maxM + 1
	r = maxI + maxM - 1
	fmt.Println(mlog.GetLineNumber(), maxI, maxM, l, r)
	return str[l>>1 : r>>1]
}

func TestManacher3(t *testing.T) {
	str := "abaxabaxabybaxabyb"

	test3(str, t)

	str = "abaxabaxabyybaxabyb"
	test3(str, t)
	//fmt.Println(mancher2(str))
	//fmt.Println(mancher3(str))

	str = "ccc"
	test3(str, t)

	str = "cccc"
	test3(str, t)

	//fmt.Println("---------")
	str = "babad"
	test3(str, t)
	//fmt.Println(mancher3(str))
	//fmt.Println(mancher2(str))

	//fmt.Println("---------")
	str = "cbbd"
	test3(str, t)
	//fmt.Println(mancher3(str))
	//fmt.Println(mancher2(str))

	//fmt.Println("---------")
	str = "ac"
	test3(str, t)
	//fmt.Println(mancher3(str))

	//fmt.Println("---------")
	str = "bb"
	test3(str, t)
	//fmt.Println(mancher3(str))

}

func test3(str string, t *testing.T) {

	want := mancher2(str)
	got := mancher3(str)
	if want != got {
		t.Fatalf("Manacher3 \r\nwant=%q, \r\ngot =%q", want, got)
	}
	/*
		got = mancher1(str)
		if want != got {
			t.Fatalf("Manacher3 want=%q, got=%q", want, got)
		}
	*/

	fmt.Println(got)
}
