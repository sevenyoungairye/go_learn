package main

import "fmt"

func init() {
	fmt.Println("切片并不存储任何数据，它只是描述了底层数组中的一段.. ")

	// 切片是数据的引用.. 更改切片的值就修改数组底层的值
	var names = [5]string{
		"jack",
		"gorgeo",
		"lucy",
		"god",
	}

	a := names[1:2]
	b := names[2:4]
	a[0] = "rose"

	fmt.Println(a, b)
	fmt.Println(names)
}
