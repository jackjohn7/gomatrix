package matrix

import (
	"fmt"
	"testing"
)

type GetTest struct {
	row      int
	col      int
	expected int
}

func TestMatrixGet(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	expectedResults := []GetTest{
		{
			row:      0,
			col:      0,
			expected: 1,
		},
		{
			row:      0,
			col:      1,
			expected: 2,
		},
		{
			row:      1,
			col:      0,
			expected: 3,
		},
		{
			row:      1,
			col:      1,
			expected: 4,
		},
		{
			row:      2,
			col:      0,
			expected: 5,
		},
		{
			row:      2,
			col:      1,
			expected: 6,
		},
	}

	m := FromData(3, 2, data)
	for _, exp := range expectedResults {
		d1, err := m.Get(exp.row, exp.col)
		if err != nil {
			t.Fatalf("Err: %s\n", err.Error())
		}

		if d1 != exp.expected {
			t.Fatalf("Mismatch: %d != %d\n", d1, exp.expected)
		}
	}
}

func TestRows(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
	m := FromData(3, 2, data)
	cols := m.Rows()
	for i, col := range cols {
		for j, v := range col {
			if v != expected[i][j] {
				t.Fatalf("Mismatch: %d != %d\n", v, expected[i][j])
			}
		}
	}
}

func TestCols(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	expected := [][]int{{1, 3, 5}, {2, 4, 6}}
	m := FromData(3, 2, data)
	cols := m.Cols()
	fmt.Printf("cols: %d\n", len(cols))
	for i, col := range cols {
		for j, v := range col {
			if v != expected[i][j] {
				t.Fatalf("Mismatch: %d != %d\n", v, expected[i][j])
			}
		}
	}
}
