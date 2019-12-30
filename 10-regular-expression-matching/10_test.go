package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

/*
aa
a*

ab
.*
*/

func TestOk(t *testing.T) {
	type (
		TT struct {
			s, p string
			e    bool
		}
	)

	ts := []*TT{
		&TT{
			"aab",
			"c*a*b",
			true,
		},
		&TT{
			"aa",
			"a*",
			true,
		},
		&TT{
			"aa",
			"a",
			false,
		},
		&TT{
			"ab",
			".*",
			true,
		},
		&TT{
			"mississippi",
			"mis*is*p*.",
			false,
		},
		&TT{
			"caaa",
			"ca*",
			true,
		},
	}

	for _, p := range ts {
		b := isMatch1(p.s, p.p)
		if b != p.e {
			logrus.Errorf("%v got:%v", p, b)
			continue
		}
		logrus.Infof("%v got:%v", p, b)
	}
}
