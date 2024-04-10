package test_0810

import (
	"math"
)

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	left := make([][]int, m)
	for i := range left {
		left[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					left[i][j] = 0
				} else {
					left[i][j] = left[i][j-1] + 1
				}
				//left[i][j] = (j == 0 ? 0 : left[i][j-1]) + 1
			}
		}
	}

	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = int(math.Min(float64(width), float64(left[k][j])))
				area = int(math.Max(float64(area), float64((i-k+1)*width)))
			}
			ret = int(math.Max(float64(ret), float64(area)))
		}
	}

	return ret
}
