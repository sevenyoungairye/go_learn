package main

import "fmt"

type Abser interface {
	// Abs 方法..
	Abs() float64
}

func init() {
	fmt.Println("======= 接口是一系列方法的集合 =======")

	v := Vertex{}

	// 在接口中, 注意方法接收的是指针
	var a Abser = &v
	fmt.Printf("a.Abs(): %v\n", a.Abs())

	myFloat := Myfloat(-88)
	a = myFloat
	fmt.Printf("a.Abs(): %v\n", a.Abs())

}

type Myfloat float64

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	fmt.Println("v: ", v)

	return v.X + v.Y
}

// 获取类型Myfloat的绝对值..
func (f Myfloat) Abs() float64 {
	if f < 0 {

		return -float64(f)
	}

	return float64(f)
}

func main() {}
