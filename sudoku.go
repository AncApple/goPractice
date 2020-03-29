package main

import (
	"errors"
	"fmt"
	"os"
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

var (
	ErrBounds     = errors.New("数字の範囲エラー")
	ErrDigit      = errors.New("無効な数字です")
	ErrFixedDigit = errors.New("はじめからある数字を置き換えることは出来ません。")
	ErrInRow      = errors.New("この行にすでに存在する数字です。")
	ErrInColumn   = errors.New("この列にすでに存在する数字です。")
	ErrInRegion   = errors.New("このブロックにはすでに存在する数字です。")
)

//数独のルール違反を判定するためのプログラム　※解くプログラムではない
func NewSudoku(digits [rows][columns]int8) *Grid {
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
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.inFixed(row, column):
		return ErrFixedDigit
	case g.inRow(row, digit):
		return ErrInRow
	case g.inColumn(column, digit):
		return ErrInColumn
	case g.inRegion(row, column, digit):
		return ErrInRegion
	}

	g[row][column].digit = digit
	return nil
}

func (g *Grid) Clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.inFixed(row, column):
		return ErrFixedDigit
	}
	g[row][column].digit = empty
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
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

func (g *Grid) inRegion(row, column int, digit int8) bool {
	//3x3のマス目で同じ数字を禁ずる
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
	s := NewSudoku([rows][columns]int8{
		{0, 9, 0, 0, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 9, 0, 0, 4, 0},
		{5, 0, 0, 0, 0, 0, 0, 0, 6},
		{6, 5, 0, 7, 3, 2, 0, 8, 4},
		{0, 0, 9, 0, 0, 0, 7, 1, 3},
		{8, 0, 0, 0, 0, 1, 0, 6, 0},
		{2, 1, 4, 8, 0, 3, 0, 0, 9},
		{0, 8, 5, 0, 7, 6, 4, 3, 2},
		{0, 7, 6, 4, 2, 9, 0, 0, 1},
	})
	//err := s.Set(0, 0, 1)
	err := s.Clear(0, 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//結果
	for _, row := range s {
		fmt.Println(row)
	}
}
