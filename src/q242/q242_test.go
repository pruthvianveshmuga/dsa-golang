package q242

import (
	"strconv"
	"testing"
)

type Case struct {
	inp []string
	out bool
}

func TestSolutions(t *testing.T) {
	solutions := []func(string, string) bool{Solution1}
	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := []Case{
				{[]string{"anagram", "nagaram"}, true},
				{[]string{"rat", "car"}, false},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if res := sol(tc.inp[0], tc.inp[1]); res != tc.out {
						t.Errorf("sol%v, testCase%v, %v: expected %v, got %v", solInd, tcInd, tc.inp, tc.out, res)
					}
				})
			}
		})
	}
}