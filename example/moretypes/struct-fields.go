package main

import "fmt"

func init() {
	fmt.Println("========== obj ==========")
	setVal()
}

func setVal() {

	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(v.X)

}

type Vertex struct {
	X int
	Y int
}
