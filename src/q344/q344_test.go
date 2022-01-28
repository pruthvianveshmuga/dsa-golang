package q344

import (
	"strconv"
	"testing"

	byteUtils "github.com/pruthvianveshmuga/dsa/utils/byte"
)

type Case struct {
	inp []byte
	out []byte
}

func TestSolutions(t *testing.T) {
	solutions := []func([]byte) {Solution1}
	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := []Case{
				{[]byte{'h','e','l','l','o'}, []byte{'o','l','l','e','h'}},
				{[]byte{'H','a','n','n','a','h'}, []byte{'h','a','n','n','a','H'}},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if sol(tc.inp); !byteUtils.DeepEqual(tc.inp, tc.out) {
						t.Errorf("sol%v, testCase%v: expected %v, got %v", solInd, tcInd, tc.out, tc.inp)
					}
				})
			}
		})
	}
}