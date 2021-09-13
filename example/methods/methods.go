package main

import (
	"fmt"
	"math"
)

func init() {

	fmt.Println("========= 方法 ========")
	// 注意!!! go中方法就是函数 只是写法可以有一点点区别
	// 方法就是一类带特殊的 接收者 参数的函数。
	// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

	v := Vertex{X: 3, Y: 4}
	fmt.Printf("v.Abs(): %v\n", v.Abs())
}

type Vertex struct {
	X, Y float64
}

// go方法的扩展写法...
func Abs(v Vertex) float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// v为接收者
func (v Vertex) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

}
