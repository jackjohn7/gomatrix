package matrix

import (
	"errors"
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

func FromData(rows int, cols int, data []int) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: data,
	}
}

func (m Matrix) Get(i int, j int) (int, error) {
	if i >= m.rows || j >= m.cols {
		return 0, errors.New("Out of bounds")
	}
	return m.data[(i*m.cols)+j], nil
}

func (m Matrix) Rows() [][]int {
	result := make([][]int, m.rows)
	for row := 0; row < m.rows; row++ {
		result[row] = make([]int, m.cols)
		for col := 0; col < m.cols; col++ {
			result[row][col] = m.data[(row*m.cols)+col]
		}
	}
	return result
}

func (m Matrix) Cols() [][]int {
	result := make([][]int, m.cols)
	for i := range result {
		result[i] = make([]int, m.rows)
	}
	for col := 0; col < m.cols; col++ {
		for row := 0; row < m.rows; row++ {
			result[col][row] = m.data[(row*m.cols)+col]
		}
	}

	return result
}
