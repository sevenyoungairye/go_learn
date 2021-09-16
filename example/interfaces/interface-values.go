package main

import "fmt"

type I interface {
	M()
}

type T0 struct {
	S string
}

type F float64

// 结构体T0实现接口I
func (t *T0) M() {
	fmt.Println(t.S)
}

// float F实现接口I
func (t F) M() {
	fmt.Println(t)
}

func describe(i I) {
	// fmt.Printf("val: %v, type: %T\n", i, i)
	fmt.Printf("(%v, %T)\n", i, i)
}

func init() {
	// 接口值保存了一个具体底层类型的具体值。
	// 接口值调用方法时会执行其底层类型的同名方法。
	fmt.Println("===== 接口的值是底层实现 =====")
	var i I = F(9)
	i.M()
	describe(i)

	i = &T0{S: "hello, hah"}
	i.M()
	describe(i)

}
