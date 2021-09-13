package main

import "fmt"

// 对比indirection.go 使用指针作为参数..
// 方法：值作为参数
func init() {

	fmt.Println("======= in val ========")

	v := Vertex{1, 1}
	// 值调用
	v.m1()
	fmt.Println(v)

	p := &v
	// 方法支持指针调用
	p.m1()
	// (*p).m1()
	fmt.Println(v)

	// 函数 参数是值就得传值, 是指针就得传指针
	f1(v)
}

// 声明方法 参数是值类型
// 复制了一份v
func (v Vertex) m1() {

	v.X = 0
	v.Y = 0
	fmt.Println("m1, ", v)
}

// 声明函数 参数为值类型
func f1(v Vertex) {
	v.X = -1
	v.Y = -1
}
