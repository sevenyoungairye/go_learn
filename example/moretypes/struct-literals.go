package main

import "fmt"

func init() {

	fmt.Println("=========literals start..========")
	literals()
}

func literals() {

	type Vertex struct {
		// 相同类型的变量声明 可简写
		X, Y int
	}

	var (
		// 隐式声明
		v1 = Vertex{} // 创建一个Vertex类型的结构体
		v2 = Vertex{1, 2}
		v3 = Vertex{Y: 3}
		p  = &Vertex{9, 9} // 创建一个*Vertex类型的结构体
	)

	fmt.Println(v1, v2, v3, p, *p)

	type student struct {
		name string
		age  int
	}
	var (
		s1 = student{name: "jack", age: 19}
		s2 = student{name: "rose", age: 20}
		s3 = student{name: "jack", age: 10}
	)
	fmt.Println(s1 == s2, s1.name == s3.name)
}
