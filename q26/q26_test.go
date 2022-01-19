package q26

import (
	"testing"
)

func TestSolution1(t *testing.T) {
	inp := []int{1,1,2}
	out := 2
	res := Solution1(inp)
	if res != out {
		t.Errorf("got %v, want %v", res, out)
	}
}

func TestSolution2(t *testing.T) {
	inp := []int{1,1,2}
	out := 2
	res := Solution2(inp)
	if res != out {
		t.Errorf("got %v, want %v", res, out)
	}
}
