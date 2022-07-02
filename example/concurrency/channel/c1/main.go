package main

import (
	"fmt"
	"time"
)

/*
通道（channel）是用来传递数据的一个数据结构。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据, 并把值赋给 v

声明：
ch := make(chan int)
默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：

*/

func main() {

	ch := make(chan int)

	var arr1 = []int{1, 2, 3}

	go sum(arr1, ch)
	time.Sleep(100 * time.Millisecond)
	go sum([]int{6, 6, 6}, ch)

	// 从通道取值
	a := <-ch
	b := <-ch
	fmt.Println("get it.. ", a, b)
}

func sum(arr []int, ch chan int) {
	var sum = 0

	for _, v := range arr {
		sum += v
	}

	// 将和放入通道
	ch <- sum
}
