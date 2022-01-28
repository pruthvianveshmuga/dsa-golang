package q48

import (
	"strconv"
	"testing"

	"github.com/pruthvianveshmuga/dsa/utils"
)

type Case struct {
	inp [][]int
	out [][]int
}

func TestSolutions(t *testing.T) {
	solutions := []func([][]int) {Solution1, Solution2}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[][]int{{1,2,3}, {4,5,6}, {7,8,9}}, [][]int{{7,4,1}, {8,5,2}, {9,6,3}}},
				{[][]int{{5,1,9,11}, {2,4,8,10}, {13,3,6,7}, {15,14,12,16}}, [][]int{{15,13,2,5}, {14,3,4,1}, {12,6,8,9}, {16,7,10,11}}},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if sol(tc.inp); !utils.DeepEqual2(tc.inp, tc.out) {
						t.Errorf("sol%v, testCase%v: expected %v, got %v", solInd, tcInd, tc.out, tc.inp)
					}
				})
			}
		})
	}
}