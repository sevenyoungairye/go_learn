package main

import (
	"fmt"
	"time"
)

// 实现内置错误接口error
func (err *MyError) Error() string {

	return err.When.String() + "\t" + err.What
}

// 获取错误信息..
func run() error {

	return &MyError{time.Now(), "sth err..."}
}

// 定义一个结构体
type MyError struct {
	When time.Time
	What string
}

func init() {

	fmt.Println("==== errors.go ====")

	if e := run(); e != nil {
		fmt.Printf("e.Error(): %v\n", e.Error())
	}

}

func main() {

}
