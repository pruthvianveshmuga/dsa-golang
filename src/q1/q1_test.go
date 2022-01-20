package q1

import (
	"strconv"
	"testing"

	"github.com/pruthvianveshmuga/dsa/utils"
)

type Case struct {
	inp1 []int
	inp2 int
	out []int
}

func TestSolutions(t *testing.T) {
	solutions := []func([]int, int) []int{Solution1, Solution2, Solution3, Solution4}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[]int{2,7,11,15}, 9, []int{0,1}},
				{[]int{3,2,4}, 6, []int{1,2}},
				{[]int{3,3}, 6, []int{0,1}},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if res := sol(tc.inp1, tc.inp2); !utils.DeepEqualWithoutOrder(res, tc.out) {
						t.Errorf("sol%v, testCase%v, %v & %v: expected %v, got %v", solInd, tcInd, tc.inp1, tc.inp2, tc.out, res)
					}
				})
			}
		})
	}
}