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
			row: 0,
			col: 1, expected: 2,
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

func TestEq(t *testing.T) {
	m1 := FromData(2, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	m2 := FromData(2, 5, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	if m1.Eq(m2) {
		t.Fatalf("Matrices should not match")
	}

	m3 := FromData(2, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	if !m1.Eq(m3) {
		t.Fatalf("Matrices should match")
	}
}

func TestAdd(t *testing.T) {
	m1 := FromData(2, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	m2 := FromData(2, 5, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	ex := FromData(2, 5, []int{12, 14, 16, 18, 20, 22, 24, 26, 28, 30})

	sum, err := m1.Add(m2)
	if err != nil {
		t.Fatalf("Err: %s\n", err.Error())
	}

	if !sum.Eq(ex) {
		t.Fatalf("Mismatch: %+v != %+v\n", sum, ex)
	}
}

func TestScale(t *testing.T) {
	m1 := FromData(2, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	ex := FromData(2, 5, []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50})
	res := m1.Scale(5)
	if !res.Eq(ex) {
		t.Fatalf("Mismatch: %+v != %+v\n", res, ex)
	}
}
