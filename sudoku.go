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

//func (g *Grid)Clear(){}

//func inRow(){}

//func inColumn(){}

//func inRegin(){}

//func inFixed(){}

func main() {
	//NewCreate実行
	s := NewCreate([rows][columns]int8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
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
