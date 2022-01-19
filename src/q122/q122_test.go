package q122

import (
	"strconv"
	"testing"
)

type Case struct {
	inp  []int
	out int
}

func TestSolutions(t *testing.T) {
	solutions := []func([]int) int{Solution1}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[]int{7,1,5,3,6,4}, 7},
				{[]int{1,2,3,4,5}, 4},
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
