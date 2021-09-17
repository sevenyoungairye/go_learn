package main

import "fmt"

func (p Person) String() string {

	return fmt.Sprintf("hello, my name is %v, and I'm %v years old..", p.Name, p.Age)
}

type Person struct {
	Name string
	Age  int
}

func init() {

	fmt.Println("==== stringer.go ====")

	fmt.Println(Person{"jack", 19})
}
