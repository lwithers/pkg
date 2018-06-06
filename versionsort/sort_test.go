package versionsort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLess(t *testing.T) {
	// ensure v1<v2 (and that !(v2<v1) — so tc cannot contain v1==v2!)
	for _, tc := range []struct {
		v1, v2 string
	}{
		{"1.1", "1.2"},
		{"1.2", "1.10"},
		{"x100", "x101"},
		{"x99", "x100"},
		{"x100x", "x101x"},
		{"x99x", "x100x"},
		{"100000000000000000000000000000000", "100000000000000000000000000000001"},
		{"abc", "xyz"},
		{"x1", "x11"},
		{"1.1.1", "1.1.2"},
		{"1.2.1", "1.2.3"},
		{"1.2.3", "1.3.1"},
	} {
		t.Run(fmt.Sprintf("%s<%s", tc.v1, tc.v2), func(t *testing.T) {
			if !Less(tc.v1, tc.v2) {
				t.Errorf("%q<%q: got false, expected true",
					tc.v1, tc.v2)
			}
			if Less(tc.v2, tc.v1) {
				t.Errorf("%q<%q: got true, expected false",
					tc.v2, tc.v1)
			}
		})
	}

	// for v1==v2, ensure that !(v1<v2)
	for _, tc := range []string{
		"1", "1.1", "x1", "x100", "x100x", "xx",
	} {
		t.Run(fmt.Sprintf("%s==%s", tc, tc), func(t *testing.T) {
			if Less(tc, tc) {
				t.Error("got true, expected false")
			}
		})
	}
}

func TestIntegerStartEnd(t *testing.T) {
	for _, tc := range []struct {
		v          string
		start      int
		expS, expE int
	}{
		{"0", 0, 0, 1},
		{"00", 0, 1, 2},
		{"0x", 0, 0, 1},
		{"00x", 0, 1, 2},
		{"1", 0, 0, 1},
		{"11", 0, 0, 2},
		{"1x", 0, 0, 1},
		{"11x", 0, 0, 2},
		{"01", 0, 1, 2},
		{"01x", 0, 1, 2},
		{"x1", 1, 1, 2},
		{"x1x", 1, 1, 2},
	} {
		t.Run(tc.v, func(t *testing.T) {
			s, e := integerStartEnd(tc.v, tc.start)
			if s != tc.expS || e != tc.expE {
				t.Errorf("got (%d,%d), expected (%d,%d)",
					s, e, tc.expS, tc.expE)
			}
		})
	}
}

func TestSort(t *testing.T) {
	exp := []string{
		"1",
		"1.1",
		"1.2",
		"2.9",
		"2.10",
		"2.11",
		"11",
		"12",
		"x",
	}

	for round := 0; round < 100; round++ {
		order := rand.Perm(len(exp))
		var perm []string
		for _, idx := range order {
			perm = append(perm, exp[idx])
		}

		Strings(perm)
		ok := true
		for i := range exp {
			if perm[i] != exp[i] {
				ok = false
			}
		}

		if !ok {
			for i := range exp {
				t.Errorf("%d: exp %q, act %q",
					i, exp[i], perm[i])
			}
			t.Fatalf("out of order on round %d", round)
		}
	}
}
