package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	var err error
	lines := strings.Split(s, "\n")
	matrix := make(Matrix, len(lines))

	for idx, line := range lines {
		cols := strings.Fields(line)
		if idx > 0 && len(cols) != len(matrix[0]) {
			return nil, errors.New("rows have unequal length")
		}
		row := make([]int, len(cols))
		for colNum, col := range cols {
			if row[colNum], err = strconv.Atoi(col); err != nil {
				return nil, errors.New("invalid int format in matrix")
			}
			matrix[idx] = row
		}
	}

	return &matrix, nil
}

// Cols returns the columns of the matrix.
func (m *Matrix) Cols() [][]int {
	mutableMatrix := *m

	if len(mutableMatrix) == 0 {
		return nil
	}
	cols := make([][]int, len(mutableMatrix[0]))

	for idx := range cols {
		col := make([]int, len(mutableMatrix))
		for j := range col {
			col[j] = mutableMatrix[j][idx]
		}
		cols[idx] = col
	}

	return cols
}

// Rows returns the rows of the matrix.
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(*m))

	for idx, row := range *m {
		rows[idx] = append(rows[idx], row...)
	}

	return rows
}

// Set sets the value of the matrix at the given row and column.
func (m *Matrix) Set(row, col, val int) bool {
	mm := *m
	if row < 0 || row >= len(mm) || col < 0 {
		return false
	}
	if cols := len(mm[0]); col >= cols {
		return false
	}
	mm[row][col] = val
	return true
}
