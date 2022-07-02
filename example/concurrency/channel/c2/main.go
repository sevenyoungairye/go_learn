package main

import (
	"fmt"
	"time"
)

func main() {

	// https://www.runoob.com/go/go-concurrent.html
	noBuffer()
	// go HasBuffer()

	time.Sleep(100 * time.Millisecond)

}

func noBuffer() {
	ch := make(chan int)
	ch <- 2
	// 无缓存的通道需要被接收，否则会造成死锁。
}

func HasBuffer() {
	// 带有缓冲大小的通道
	ch := make(chan int, 2)

	ch <- 2
	ch <- 1

	// 阻塞
	// ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
