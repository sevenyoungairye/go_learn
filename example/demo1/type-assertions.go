package main

import "fmt"

func init() {

	fmt.Println("===== type-assertion.go =====")

	// i 为接口
	// t := i.(T)
	var i interface {
	} = "hello"

	assert, ok := i.(string)

	if ok {
		fmt.Println("断言成功.. 返回底层值: ", assert, ok)
	}

	f, ok := i.(float64)
	fmt.Println("fail.. ", f, ok) // 0, false

	// panic: interface conversion: interface {} is string, not float64
	// f = i.(float64)
	// fmt.Println(f)
}
