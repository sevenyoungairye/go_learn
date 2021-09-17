package main

import "fmt"

func init() {

	fmt.Println("====  type-switches.go ====")
	do("hello")

	do(1)

	do(true)
}

// 断言: 类型选择
func do(i interface{}) {

	switch v := i.(type) {
	case string:
		fmt.Printf("value: %q, type: %T\n", v, v)
	case int:
		fmt.Printf("value: %v, type: %T\n", v, v)
	default:
		fmt.Println("nothing here.. default.. ")
	}
}
