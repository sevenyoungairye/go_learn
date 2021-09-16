package main

import "fmt"

// 隐式实现了拥有M的接口..
func (t1 *T1) M() {
	// 接收者可能是 nil
	if nil == t1 {
		fmt.Println("t1 is ", t1)
		return
	}
	fmt.Println("Value of type T1 is:", t1.S)
}

type T1 struct {
	S string
}

func init() {

	fmt.Println("======= 接口底层实现是nil的情况 =======")
	var i I
	var t1 *T1
	i = t1
	i.M()

	t1 = &T1{S: "sure"}
	i = t1
	i.M()
}
