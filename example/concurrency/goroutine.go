package main

import (
	"fmt"
	"time"
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
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s+"\t", i)
	}
}
