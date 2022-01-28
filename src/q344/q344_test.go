package q344

import (
	"strconv"
	"testing"
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
					if sol(tc.inp); !DeepEqual(tc.inp, tc.out) {
						t.Errorf("sol%v, testCase%v: expected %v, got %v", solInd, tcInd, tc.out, tc.inp)
					}
				})
			}
		})
	}
}

func DeepEqual(a, b []byte) bool {
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

func Swap(a *byte, b *byte) {
	temp := *a
	*a = *b
	*b = temp
}