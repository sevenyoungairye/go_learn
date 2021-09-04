package main

import (
	"fmt"
)

func init() {

}

func main() {
	fmt.Println("==========moretypes start==========")

	goPointer()

	basicStruct()
}

func basicStruct() {

	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})
}

// 指针: 间接引用or重定向
func goPointer() {
	// go 指针.. 指针可以理解为 一个中间工具
	//1. * 代表指针指向的底层值
	var p *int
	fmt.Println("指针的零值", p, p == nil)

	i, k := 22, 88
	//2. 使用 & 生成指针, 指向i
	j := &i
	//3. 通过指针读取i的值
	fmt.Println("j=", j, "使用*获取指针值", *j)
	fmt.Println(i)

	//4. 设置指针的值
	*j = 11
	fmt.Println(j, *j, i)

	// 指向k
	j = &k
	*j = *j / i
	fmt.Println(j, *j, i, k)

}
