package main

import (
	"fmt"
	"math"
)

func init() {

	fmt.Println("============= range语法糖.. ============")

	var pow []int = make([]int, 10)

	// 下标..
	for index := range pow {
		pow[index] = int(math.Pow(2, float64(index)))
	}

	// 值..
	for _, val := range pow {
		fmt.Printf("%d ", val)
	}
	fmt.Println("")
}
