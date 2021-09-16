package main

import "fmt"

func init() {
	// 指定了零个方法的接口值被称为 *空接口：*
	// 空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

	fmt.Println("===== 空接口 =====")
	var a interface{}

	describeInter(a)

	a = T{S: "hah"}
	describeInter(a)

}

func describeInter(i interface{}) {

	fmt.Printf("(%v, %T)\n", i, i)
}
