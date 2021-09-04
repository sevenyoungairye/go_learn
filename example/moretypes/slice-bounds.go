package main

import (
	"fmt"
)

func init() {

	fmt.Println("======== 切边边界值 =======")

	arr := [10]int{2, 3, 5, 7, 11}
	var lenArr = len(arr)
	s0 := arr[0:]
	s1 := arr[0:lenArr]
	s2 := arr[:lenArr]
	s3 := arr[:]

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}
