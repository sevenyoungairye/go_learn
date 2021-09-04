package main

import (
	"fmt"
)

func init() {
	fmt.Println("========= 切片的零值.. =========")
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}
