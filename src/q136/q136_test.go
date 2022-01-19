package q136

import (
	"strconv"
	"testing"
)

type Case struct {
	inp  []int
	out int
}

func TestSolutions(t *testing.T) {
	solutions := []func([]int) int{Solution1, Solution2, Solution3}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[]int{2,2,1}, 1},
				{[]int{4,1,2,1,2}, 4},
				{[]int{1}, 1},
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
