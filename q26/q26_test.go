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
		// Cannot move this to the top of the loop, because the test cases are not deep-copied.
		testCases := []Case{
			{[]int{1,1,2}, 2},
			{[]int{0,0,1,1,1,2,2,3,3,4}, 5},
		}
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			for tcInd, tc := range testCases {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if res := sol(tc.inp); res != tc.out {
						t.Errorf("sol%v, testCase%v, %v: expected %v, got %v", solInd, tcInd, tc.inp, tc.out, res)
					}
				})
			}
		})
	}
}
