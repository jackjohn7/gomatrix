package matrix

import (
	"fmt"
)

type Matrix struct {
	rows int
	cols int
	data []int
}

func New(rows int, cols int) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]int, rows*cols),
	}
}

func (m Matrix) Get(i int, j int) int {
	return m.data[(i*m.cols)+(j*m.rows)]
}

//func (m Matrix) Cols() []int {
//	result := make([][]int, m.cols*m.rows)
//	for i := 0; i < m.rows; i++ {
//		for j := 0; j < m.cols; j++ {
//			result[i][j] = m.data[]
//		}
//	}
//
//	return result
//}
