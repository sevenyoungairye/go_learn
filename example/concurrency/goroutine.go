package main

import (
	"fmt"
	"runtime"
)

func init() {}

func main() {

	// 通过go关键字启动gorutine
	go say("world")
	// 当前gorutines执行
	say("hello")

}

func say(s string) {
	for i := 0; i < 5; i++ {
		// 让其它线程先跑
		runtime.Gosched()
		fmt.Println(s + "\t")
	}
}
