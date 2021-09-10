package main

import "fmt"

func init() {

	fmt.Println("======= map crud ========")

	f := func() {
		var m map[string]int = map[string]int{
			"Answer": -1,
		}

		fmt.Println(m["Answer"])

		// 修改键
		m["Answer"] = 99
		fmt.Println(m)

		// 删除键
		delete(m, "Answer")
		fmt.Println("del map.. ", m)

		// 判断键值是否存在
		val, ok := m["Answer"]
		fmt.Println("val ", val, " is present? ", ok)
	}

	f()

}
