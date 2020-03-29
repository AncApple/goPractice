package main

import "fmt"

func main() {
	var command = "go east"

	//もしコマンドが"g east"とひとしければ
	if command == "go east" {
		fmt.Println("きみは、更に山を登る。")
	} else if command == "go inside" {
		fmt.Println("きみは洞窟に入り、そこで一生を過ごす。")
	} else {
		fmt.Println("なんだか、よくわからない。")
	}
}
