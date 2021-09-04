package main

import "fmt"

func init() {
	fmt.Println("数组的切片声明.. 可不用声明长度..")

	var arr = []string{"ha", "ha"}

	boolArr := []bool{true, false, true, true}

	objArr := []struct {
		str string
		b   bool
	}{
		{"jack", true},
		{"rose", true},
		{"lucy", false},
	}

	fmt.Println(arr, "\n", boolArr, "\n", objArr)
}
