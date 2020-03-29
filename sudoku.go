package main

import (
	"fmt"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

//マス目Cell
type Cell struct {
	digit int8
	fired bool
}

type Grid [rows][columns]Cell

func NewCreate(digits [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fired = true
			}
		}
	}
	return &grid
}

func (g *Grid) Set(row, column int, digit int8) error {
	g[row][column].digit = digit
	return nil
}

func (g *Grid) Clear(row, column int) error {
	g[row][column].digit = empty
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inRegin(row, column int, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func (g *Grid) inFixed(row, column int) bool {
	return g[row][column].fired
}

func main() {
	//NewCreate実行
	s := NewCreate([rows][columns]int8{
		{1, 1, 0, 0, 0, 0, 0, 0, 9},
		{2, 3, 0, 0, 0, 0, 0, 0, 0},
		{2, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	//	err := s.Set(1, 1, 4)
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//結果
	for _, row := range s {
		fmt.Println(row)
	}
}
