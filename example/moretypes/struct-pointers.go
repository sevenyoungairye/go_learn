package main

import "fmt"

func init() {
	fmt.Println("========== 结构体指针.. ==========")
	structPointer()
}

func structPointer() {
	var v Vertex = Vertex{1, 2}
	// 指向v
	p := &v
	p.X = 9
	// *p = Vertex{6, 6}
	fmt.Println(v)
}
