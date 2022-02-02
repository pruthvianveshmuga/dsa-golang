package q387

import (
	"strconv"
	"testing"
)

type Case struct {
	inp string
	out int
}

func TestSolutions(t *testing.T) {
	solutions := []func(string) int{Solution1}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{"leetcode", 0},
				{"loveleetcode", 2},
				{"aabb", -1},
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