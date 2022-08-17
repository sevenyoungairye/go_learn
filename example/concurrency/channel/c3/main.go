package main

import "fmt"

func main() {

	ch := make(chan int, 10)

	// 放10次数据到通道内, 缓冲区的长度为10
	go demo(cap(ch), ch)

	for v := range ch {
		// range 函数接收通道的值
		fmt.Println(v)
	}
}

func demo(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	// 关闭通道
	close(ch)
}
