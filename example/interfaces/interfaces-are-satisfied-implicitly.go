package main

import "fmt"

func init() {
	fmt.Println("====== 接口默认隐式实现..  =====")
	var i ImplicityInter = T{S: "hello world.."}
	i.m()
}

type ImplicityInter interface {
	m()
}

// 结构体T
type T struct {
	S string
}

// 结构体T的方法 T实现了ImplicityInter接口
// 隐式的实现了接口.. 不用implements关键字..
func (t T) m() {
	fmt.Println(t.S)
}
