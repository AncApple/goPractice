package main

inport {
	"errors"
	"fmt"
	"os"
}

const {
	rows , columns = 9,9
	empty          = 0
}

//マス目->Cell
type Cell struct {
	digit int8
	fired bool
}

type grid [rows][columns]Cell


//数独グリッドの作成
func NewCreate(digits [rows][columns]int8) *grid{

}


func (g *gird) Set(row, column int , digit int8) {
}

//
func Clear{

}

func inRow(){

}

func inColumn(){

}

func inRegin(){

}

func inFixed(){

}

func main(){
	//NewCreate実行
}
