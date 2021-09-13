package main

import (
	"fmt"
	"math"
)

// 非结构体类型声明方法..
// var f float64 = 3.0
type MyFloat float64

// 接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法
func (f MyFloat) myAbs() float64 {

	if f < 0 {
		return (-float64(f))
	}

	return float64(f)
}

func init() {

	fmt.Println("======= methods continued.. =======")
	f := MyFloat(math.Sqrt2)
	fmt.Printf("f.myAbs(): %v\n", f.myAbs())
}
