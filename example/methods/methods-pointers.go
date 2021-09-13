package main

import "fmt"

func init() {
	fmt.Println("======== 指针复习 ========")
	// 声明指针a
	var a *int
	b := 1
	// 获取到b的内存地址..
	a = &b
	fmt.Println(a, &b, " 获取到指针变量的值 ", *a)

	mehtodPointer()
	fmt.Println("== method/func split.. ==")
	funcPointer()
}

func mehtodPointer() {
	v := Vertex{1, 1}
	v.copyDemo()
	fmt.Println("by copyDemo, value wasn't changed... ", v)

	// 参数是指针, 方法能为值v
	v.changeOrg()
	fmt.Println("by changeOrg... ", v)

	// 当参数是指针, 方法能接收指针
	p := &Vertex{}
	p.changeOrg()

	// 而函数参数是指针 只能接收指针..
}

// see methods.go
// Vertex 已经声明, 结构体的方法..
func (v Vertex) copyDemo() {
	v.X, v.Y = 9.0, 9.0
	fmt.Println("in copyDemo ", v)
}

// 结构体的方法
func (v *Vertex) changeOrg() {
	v.X, v.Y = 3.0, 3.0
	fmt.Println("in changeOrg ", *v)
}

// 函数指针
func funcPointer() {
	type Vertex struct {
		X, Y float64
	}

	// 传递了一个副本v..
	f := func(v Vertex) float64 {
		v.X = 3
		v.Y = 3
		return v.X + v.Y
	}

	// 传递指针会改变底层值..
	f1 := func(v *Vertex) {
		v.X = 9
		v.Y = 9
	}

	v := Vertex{1, 2}
	fmt.Printf("f(v): %v\n", f(v))

	fmt.Println("v is: ", v)

	// 指针.. 会改变init中的v
	f1(&v)

	fmt.Println("v is: ", v)
}
