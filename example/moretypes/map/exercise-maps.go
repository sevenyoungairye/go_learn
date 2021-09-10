package main

import (
	"fmt"
	"strings"
)

func init() {

	fmt.Println("======== =======")

	// 统计每个单词出现的个数..
	f := func(s string) map[string]int {

		var res = make(map[string]int)

		keys := strings.Fields(s)

		for i, key := range keys {

			// 获取到值
			tmp := res[key]

			if key == keys[i] {
				tmp++
			}
			res[key] = tmp

		}

		return res
	}

	m := f("  foo bar foo  baz   ")
	fmt.Println(m)

	fmt.Printf("Fields are: %q \n", strings.Fields("  foo bar  baz   "))

}
