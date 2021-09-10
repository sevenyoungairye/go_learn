package main

import "fmt"

func init() {

	fmt.Println("================ range way.. ==============")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	// i为index, v为每个下标对应的值
	for i, v := range pow {
		// 2的n次方..
		fmt.Printf("2**%d = %d \n", i, v)
	}
}
