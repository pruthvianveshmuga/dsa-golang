package utils

import "sort"

type Case struct {
	inp  []int
	out int
}

func DeepCopyTestCases(testCases []Case) []Case {
	tcs := make([]Case, len(testCases))
	for tcInd, tc := range testCases {
		inp := make([]int, len(tc.inp))
		copy(inp, tc.inp)
		out := tc.out
		tcs[tcInd] = Case{inp, out}
	}
	return tcs
}

func DeepEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func DeepEqualWithoutOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

