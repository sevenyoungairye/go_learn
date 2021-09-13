package main

import "fmt"

func init() {
	fmt.Println("====== 方法与指针重定向 ======")

	v := Vertex{99, 99}

	// 接收者是指针, 方法能用值类型调用
	// 由于 changeOrg 方法有一个指针接收者，为方便起见，Go 会将语句 v.changeOrg(5) 解释为 (&v).changeOrg(5)。
	v.changeOrg()
	// (&v).changeOrg()
	fmt.Println("by changeOrg... ", v)

	// 参数是指针, 方法接收指针
	p := &Vertex{}
	p.changeOrg()

	// 而函数参数是指针 只能接收指针..
}
