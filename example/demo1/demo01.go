package main

import "fmt"

func init() {}

type I interface {
	// 方法列表..
	M()
}

type T struct {
	S string
}

func (t T) M() {

	fmt.Println(t.S)
}

func main() {
	fmt.Println("===== demo01.go =====")
	var i I = T{S: "hello .."}
	i.M()

	// 为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：
	// 其底层值以及一个报告断言是否成功的布尔值。
	t, ok := i.(T)

	// 若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。
	// 否则，ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生恐慌。
	fmt.Println(t, ok)
}
