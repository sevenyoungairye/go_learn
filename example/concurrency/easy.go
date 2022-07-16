package main

import "fmt"

func init() {
	// fmt.Println("go")

	go fmt.Println("hello world..")

	go fmt.Println("run sth..")

	fmt.Println("============")

	go run("after main run")

	go run("demo")

	run("main first..")
}

func run(s string) {
	for i := 0; i < 5; i++ {

		fmt.Println(s+"\t", i)
	}
}
