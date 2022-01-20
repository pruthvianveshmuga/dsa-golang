package q36

import (
	"strconv"
	"testing"
)

type Case struct {
	inp [][]string
	out bool
}

func TestSolutions(t *testing.T) {
	solutions := []func([][]byte) bool{Solution1}

	for solInd, sol := range solutions {
		t.Run("s"+strconv.Itoa(solInd+1), func(t *testing.T) {
			tcs := 	[]Case{
				{[][]string{
					{"5","3",".",".","7",".",".",".","."}, 
					{"6",".",".","1","9","5",".",".","."}, 
					{".","9","8",".",".",".",".","6","."}, 
					{"8",".",".",".","6",".",".",".","3"}, 
					{"4",".",".","8",".","3",".",".","1"}, 
					{"7",".",".",".","2",".",".",".","6"}, 
					{".","6",".",".",".",".","2","8","."}, 
					{".",".",".","4","1","9",".",".","5"}, 
					{".",".",".",".","8",".",".","7","9"},
				}, true},
				{[][]string{
					{"8","3",".",".","7",".",".",".","."},
					{"6",".",".","1","9","5",".",".","."},
					{".","9","8",".",".",".",".","6","."},
					{"8",".",".",".","6",".",".",".","3"},
					{"4",".",".","8",".","3",".",".","1"},
					{"7",".",".",".","2",".",".",".","6"},
					{".","6",".",".",".",".","2","8","."},
					{".",".",".","4","1","9",".",".","5"},
					{".",".",".",".","8",".",".","7","9"},
				}, false},
			}
			for tcInd, tc := range tcs {
				t.Run("c"+strconv.Itoa(tcInd+1), func(t *testing.T) {
					if res := sol(convertInput(tc.inp)); res != tc.out {
						t.Errorf("sol%v, testCase%v, %v: expected %v, got %v", solInd, tcInd, tc.inp, tc.out, res)
					}
				})
			}
		})
	}
}

func convertInput(inp [][]string) [][]byte {
	out := [][]byte{}
	for i := 0; i < len(inp); i++ {
		row := []byte{}
		for j := 0; j < len(inp[i]); j++ {
			row = append(row, byte(inp[i][j][0]))
		}
		out = append(out, row)
	}
	return out
}