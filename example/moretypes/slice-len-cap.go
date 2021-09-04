package main

import (
	"fmt"
)

func init() {

	fmt.Println("获取切片的长度和容量.. ")

	s := []int{1, 3, 5, 7, 11, 13}

	s = s[:0]
	fmt.Println(s) // len: 0, cap: 6
	printSlice(s)  // []

	s = s[1:3]
	fmt.Print(s)
	printSlice(s)

	ext := s[0:5]
	ext[0] = 0
	fmt.Println(ext, "~~", s[0:5])
}

func printSlice(s []int) {
	// cap是整个切片的容量 len是当前切片的元素个数..
	fmt.Printf("len: %d, cap: %d \n", len(s), cap(s))
}
