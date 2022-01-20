package q283

import (
	"strconv"
	"testing"

	"github.com/pruthvianveshmuga/dsa/utils"
)

type Case struct {
	inp []int
	out []int
}

func TestSolutions(t *testing.T) {
	solutions := []func([]int) {Solution1}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[]int{0,1,0,3,12}, []int{1,3,12,0,0}},
				{[]int{0}, []int{0}},
				{[]int{1,2}, []int{1,2}},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if sol(tc.inp); !utils.DeepEqual(tc.inp, tc.out) {
						t.Errorf("sol%v, testCase%v: expected %v, got %v", solInd, tcInd, tc.out, tc.inp)
					}
				})
			}
		})
	}
}