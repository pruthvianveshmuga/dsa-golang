package q48

import "github.com/pruthvianveshmuga/dsa/utils"

func flipDiagonally(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			utils.Swap(&matrix[i][j], &matrix[n-j-1][n-i-1])
		}
	}
}

func flipVertically(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			utils.Swap(&matrix[i][j], &matrix[n-i-1][j])
		}
	}
}

func Solution1(matrix [][]int) {
	flipDiagonally(matrix)
	flipVertically(matrix)
}