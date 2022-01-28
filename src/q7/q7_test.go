package q7

import (
	"strconv"
	"testing"
)

type Case struct {
	inp int
	out int
}

func TestSolutions(t *testing.T) {
	solutions := []func(int) int{Solution1, Solution2}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{-2147483648, 0},
				{1534236469, 0},
				{-123, -321},
				{123, 321},
				{120, 21},
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