package main

import (
	"fmt"
)

func init() {
	fmt.Println("=========== slice demo ===========")
	initSlice()
	s0 := []int{1, 2, 3}
	b := s0[:]
	b = appendSlice(b, 4)
	b = appendSlice(b, 5)
	fmt.Println("追加后的值.. ", b)

	copySlice()

	var s = []int{1, 3, 5, 7, 11}
	// delEle(4, s)
	insertEle(2, s, 99)

	insertNewSlice(0, []int{1, 2, 3}, []int{9, 8})
}

/**
在索引i的位置插入长度为 j 的新切片
*/
func insertNewSlice(i int, source []int, newSlice []int) {

	var lenSource int = len(source)

	var res []int = make([]int, 0)

	for index := 0; index < lenSource; index++ {
		if i == index {
			var left []int = source[:i]
			res = append(append(res, left...), newSlice...)

			var right []int = source[i:lenSource]
			res = append(res, right...)
		}
	}

	fmt.Println(res)
}

/*
在索引i的位置插入元素
*/
func insertEle(i int, s []int, ele int) {

	if s == nil {
		return
	}

	var lenS int = len(s)
	res := make([]int, i)
	for index := range s {
		if index == i {
			// 1. get left, note: don't change the left, it will be influence the basic slice..
			var left []int = s[0:i]
			copy(res, left)
			// 2. append the element ..
			res = append(res, ele)
			// 3. append before 2 steps ..
			var right []int = s[i:lenS]
			res = append(res, right...)
			// for j := 0; j < len(right); j++ {
			// 	res = append(res, right[j])
			// }
		}
	}

	fmt.Println(res)
}

/*
删除处于索引i的元素
*/
func delEle(i int, s []int) {

	l := len(s)

	if i < 0 || i >= l || s == nil || l <= 0 {
		return
	}

	fmt.Println("原数组 ", s)

	var res = make([]int, l-1)

	for index := range s {

		if i == index {
			var left = s[0:index]
			fmt.Println("left ", left)
			copy(res, left)

			var right []int = s[index+1 : l]
			fmt.Println("right ", right)
			for i := 0; i < len(right); i++ {
				left = append(left, right[i])
			}

			res = left
		}
	}
	fmt.Println("移除下标", i, res)
}

func copySlice() {
	a := []int{1, 2, 3}
	b := make([]int, 1)
	// 将a切片的值复制到b切片.. 下标对应..
	// b长度大的, 默认零值
	// b长度为0, 不会报错
	// b长度小于a, 有几个复制几个..
	copy(b, a)
	fmt.Println("a,", a, " b,", b)
}

func initSlice() {
	var s1 = []int{}
	var s2 = []int{1, 3, 5, 7, 11}
	var s3 = make([]int, 5)
	var s4 = make([]int, 5, 10)
	fmt.Printf("s1 %d, s2 %d, s3 %d, s4 %d \n", s1, s2, s3, s4)
}

// 切片追加操作..
func appendSlice(s []int, val int) (res []int) {
	return append(s, val)
}
