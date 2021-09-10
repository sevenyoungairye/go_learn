package main

import "fmt"

func init() {

	fmt.Println("======== 快速初始化map集合 =======")

	type Vertex struct {
		Lat, Long float64
	}

	f := func() {
		var m = map[string]Vertex{
			"Google": {23, 120},
			"Oracle": {12, 34},
		}
		fmt.Println(m)
	}

	f()

}
