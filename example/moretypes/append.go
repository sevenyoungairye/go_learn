package main

import "fmt"

func init() {
	fmt.Println("========== append sth..  ===========")

	var s []int
	// 1. append 会自动扩容和增长
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	// 可添加多个..
	s = append(s, 1, 2, 3, 4)
	printSlice(s)

	// 可添加切片
	s = append(s, []int{1, 2, 3}...)
	printSlice(s)

	// 添加切片扩容方式：按原来的容量增加一倍
	s = append(s, []int{3, 4, 9, 9, 10}...)
	printSlice(s)
}
