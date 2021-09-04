package main

import (
	"fmt"
)

func init() {

	fmt.Println("======array go========")

	var arr [10]int
	arr[0] = 1
	arr[1] = 2
	// 未赋值的为对应的零值
	fmt.Println(arr)

	var brr = [2]string{"tom"}
	fmt.Println(brr)

	strs := [10]string{"jack", "rose", "handsome lee.."}
	// strs[0] = "hah"
	fmt.Println(strs, len(strs), strs[3] == "")

	var vertexArr [10]Vertex
	vertexArr[0] = Vertex{10, 20}
	fmt.Println("对象的默认值:", vertexArr[1]) // {0 0}
}
