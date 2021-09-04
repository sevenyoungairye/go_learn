package main

import "fmt"

func init() {
	fmt.Println("===========slices start=========")

	var arr = [10]int{3, 4, 5, 9}
	// 包前不包后..
	brr := arr[1:2]
	fmt.Println(brr)
}
