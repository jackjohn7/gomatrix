package matrix

import (
	"errors"
)

type Matrix struct {
	rows int
	cols int
	data []int
}

/*
Produce an empty matrix (values are zeroed)
*/
func New(rows int, cols int) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]int, rows*cols),
	}
}

/*
Given some data, produce a Matrix.
*/
func FromData(rows int, cols int, data []int) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: data,
	}
}

/*
Returns the specified value in the matrix or an error when
accessing out of bounds.
*/
func (m Matrix) Get(i int, j int) (int, error) {
	if i >= m.rows || j >= m.cols {
		return 0, errors.New("Out of bounds")
	}
	return m.data[(i*m.cols)+j], nil
}

/*
Returns the row vectors of the Matrix
*/
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

/*
Returns the column vectors of the Matrix.
*/
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

/*
Compares two matrices. If they are differing sizes, the return is false,
this may be changed to return an error later. Otherwise, compare each
element.
*/
func (m Matrix) Eq(other *Matrix) bool {
	if m.rows == other.rows && m.cols == other.cols {
		for i := 0; i < m.cols*m.rows; i++ {
			if m.data[i] != other.data[i] {
				return false
			}
		}
		return true
	}
	return false
}

/*
Produces the sum of two matrices in a new Matrix.
*/
func (m Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.cols != other.cols || m.rows != other.rows {
		return nil, errors.New("mismatching dimensions")
	}

	data := make([]int, m.rows*m.cols)
	for i := 0; i < m.rows*m.cols; i++ {
		data[i] = m.data[i] + other.data[i]
	}
	return FromData(m.rows, m.cols, data), nil
}

/*
Scales a matrix by a scalar value and returns a new Matrix
without modifying the previous.
*/
func (m Matrix) Scale(scalar int) *Matrix {
	data := make([]int, m.rows*m.cols)
	for i := 0; i < m.rows*m.cols; i++ {
		data[i] = m.data[i] * scalar
	}

	return FromData(m.rows, m.cols, data)
}
