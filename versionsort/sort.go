/*
Package versionsort implements a version-aware lexicographical sort.
For example, it knows that "1.2" is in fact less than "1.10".
*/
package versionsort

import "sort"

// Versions can be sorted.
type Versions []string

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Less(i, j int) bool {
	return Less(v[i], v[j])
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// Sort an array of strings using version sort.
func Strings(versions []string) {
	v := Versions(versions)
	sort.Sort(v)
}

// Less returns true if v1 < v2, else false.
func Less(v1, v2 string) bool {
	if v1 == v2 {
		return false
	}

	// vStart is set when:
	//  - strings are equal up to ‘pos’
	//  - v1[pos-1] is not an integer
	//  - v1[pos] is an integer
	// i.e. it should point to the start of any version number
	vStart := -1

	smallestLen := len(v1)
	if len(v2) < smallestLen {
		smallestLen = len(v2)
	}

	// iterate over each byte, breaking if we differ or reach the end of
	// the string — and keep vStart updated
	var pos int
	for pos = 0; pos < smallestLen; pos++ {
		if v1[pos] != v2[pos] {
			break
		}

		// strings are equal thus far

		if vStart == -1 {
			// if we are not inside an integer, check for one starting
			if v1[pos] >= '0' && v1[pos] <= '9' {
				vStart = pos
			}

		} else {
			// if we are inside an integer, check for it finishing
			if v1[pos] < '0' || v1[pos] > '9' {
				vStart = -1
			}
		}
	}

	// test for non-version-number differences
	switch {
	case pos == len(v1):
		// v1 is a prefix of v2
		return true
	case pos == len(v2):
		// v2 is a prefix of v1
		return false
	case vStart == -1 && (v1[pos] < '0' || v1[pos] > '9' || v2[pos] < '0' || v2[pos] > '9'):
		// a non-numeric difference occurred, so fall back to standard
		// lexicographical comparison
		return v1 < v2
	}

	// extract position of integer fragment
	v1s, v1e := integerStartEnd(v1, pos)
	v2s, v2e := integerStartEnd(v2, pos)

	// test for the smaller version number
	switch {
	// if version in v1 has a larger magnitude, it is necessarily greater
	case (v1e - v1s) > (v2e - v2s):
		return false

	// similarly if v2 has larger magnitude, v1 is the smallest
	case (v1e - v1s) < (v2e - v2s):
		return true

	// otherwise, same magnitude, so a lexicographical comparison of the
	// numeric fragment will suffice
	default:
		return v1[v1s:v1e] < v2[v2s:v2e]
	}
}

// integerStartEnd finds the start and end positions of an integer whose first
// byte lies at p. The start position may not equal p if there are leading
// zeroes.
func integerStartEnd(s string, p int) (start, end int) {
	start = p
	for end = p; end < len(s); end++ {
		if s[end] < '0' || s[end] > '9' {
			break
		}
	}
	for start < end-1 && s[start] == '0' {
		start++
	}
	return
}
