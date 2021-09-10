package main

import "fmt"

func init() {
	fmt.Println("========== 函数闭包.. ==========")

	f := adder()
	// fmt.Printf("f(10): %v\n", f(10))

	var res int
	for i := 0; i < 10; i++ {
		res = f(i)
	}
	// 45
	fmt.Println(res)
}

// 声明： 函数 + 匿名函数
func adder() func(int) int {

	var sum int = 0
	// 调用1次
	fmt.Println("sum... ", sum)
	// 返回匿名函数..
	return func(i int) int {

		sum += i

		return sum
	}
}
