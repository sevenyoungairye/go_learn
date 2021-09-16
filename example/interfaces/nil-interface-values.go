package main

import "fmt"

func init() {

	fmt.Println("===== nil =====")
	var i I
	describe(i)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// i.M()
}
