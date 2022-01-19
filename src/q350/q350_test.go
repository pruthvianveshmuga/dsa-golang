package q350

import (
	"strconv"
	"testing"
)

type Case struct {
	inp [][]int
	out []int
}

func TestSolutions(t *testing.T) {
	solutions := []func([]int, []int) []int{Solution1, Solution2}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[][]int{{1,2,2,1}, {2,2}}, []int{2,2}},
				{[][]int{{4,9,5}, {9,4,9,8,4}}, []int{4,9}},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if res := sol(tc.inp[0], tc.inp[1]); !deepEqual(res, tc.out) {
						t.Errorf("sol%v, testCase%v, %v: expected %v, got %v", solInd, tcInd, tc.inp, tc.out, res)
					}
				})
			}
		})
	}
}

func deepEqual(a, b []int) bool {
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