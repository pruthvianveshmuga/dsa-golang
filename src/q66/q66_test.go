package q66

import (
	"strconv"
	"testing"

	intUtils "github.com/pruthvianveshmuga/dsa/utils/int"
)

type Case struct {
	inp []int
	out []int
}

func TestSolutions(t *testing.T) {
	solutions := []func([]int) []int{Solution1}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[]int{1,2,3}, []int{1,2,4}},
				{[]int{4,3,2,1}, []int{4,3,2,2}},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if res := sol(tc.inp); !intUtils.DeepEqual(res, tc.out) {
						t.Errorf("sol%v, testCase%v, %v: expected %v, got %v", solInd, tcInd, tc.inp, tc.out, res)
					}
				})
			}
		})
	}
}