package main

import "fmt"

func init() {

	fmt.Println("========== map mappings... ==========")

	// 匿名函数..
	f := func() {
		type Vertex struct {
			Lat, Long float64
		}

		// 声明map, key是string, value是Vertex
		// make 也可用来构建map集合
		var m map[string]Vertex = make(map[string]Vertex)

		m["Aisa/ShangHai"] = Vertex{121.473658, 31.230378}

		fmt.Println(m["Aisa/ShangHai"])
	}

	f()
}

func main() {

}
