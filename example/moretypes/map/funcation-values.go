package main

import (
	"fmt"
	"math"
)

// 以函数作为参数值..
func compute(fn func(x, y float64) float64) float64 {

	// 给定被传入函数 参数
	return fn(3, 4)
}
func init() {
	fmt.Println("============ 以函数作为参数值.. ============")

	// 原意：假设 计算斜边用
	hypot := func(x, y float64) float64 {

		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("hypot(5, 12): %v\n", hypot(5, 12))
	fmt.Printf("compute(hypot): %v\n", compute(hypot))

	// 3的立方
	powVal := compute(math.Pow)
	fmt.Println(powVal)
}
