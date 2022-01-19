package q26

import (
	"strconv"
	"testing"
)

type Case struct {
	inp  []int
	out int
}

func TestSolutions(t *testing.T) {
	solutions := []func([]int) int{Solution1, Solution2}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			// Use deepCopyTestCases() if you move the testcases to the top level.
			tcs := 	[]Case{
				{[]int{1,1,2}, 2},
				{[]int{0,0,1,1,1,2,2,3,3,4}, 5},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if res := sol(tc.inp); res != tc.out {
						t.Errorf("sol%v, testCase%v, %v: expected %v, got %v", solInd, tcInd, tc.inp, tc.out, res)
					}
				})
			}
		})
	}
}

func deepCopyTestCases(testCases []Case) []Case {
	tcs := make([]Case, len(testCases))
	for tcInd, tc := range testCases {
		inp := make([]int, len(tc.inp))
		copy(inp, tc.inp)
		out := tc.out
		tcs[tcInd] = Case{inp, out}
	}
	return tcs
}
