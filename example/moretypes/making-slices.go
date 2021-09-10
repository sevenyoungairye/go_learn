package main

import "fmt"

func init() {
	fmt.Println("====== 使用make创建切片.. 动态创建数组 ======")

	// 1. 创建指定长度的切片 每个元素的值为该类型的零值
	a := make([]int, 5)
	fmt.Println(a, len(a), cap(a))

	// 2. 创建指定容量, 数组长度为0, 容量为5
	b := make([]int, 0, 5)
	fmt.Println(b, len(b), cap(b))

	// 截取长度为0的切片, 注意数组越界问题.. 比如 c := b[1:]
	c := b[:3]
	fmt.Println(c, len(c), cap(c))

	d := c[2:]
	fmt.Println(d, len(d), cap(d))

	f := make([]int, 8)
	for i := 0; i < cap(f); i++ {
		f[i] = i
	}
	fmt.Println("f slice...", f)

	// 后面有值 容量不变
	// len: 4, cap: 8
	f1 := f[:4]
	// 前面有值 改变容量
	// len: 5, cap: 5
	f2 := f[3:]

	printSlice(f1)
	printSlice(f2)

}
