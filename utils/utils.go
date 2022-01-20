package utils

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