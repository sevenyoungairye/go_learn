package main

import (
	"fmt"
	"time"
)

var ch chan int

func main() {
	ch = make(chan int)
	go produceChan(ch)
	go consume(ch)

	time.Sleep(time.Second)

	fmt.Println("buffer channel")

	bufferChan()
}

func produceChan(ch chan int) {
	fmt.Println("发送了消息, 生产中的chan", ch)
	ch <- 1 // 给管道发送消息
}

func consume(ch chan int) {
	fmt.Println("接收到了消息, 消费中的chan", ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func bufferChan() {
	var buffChan chan int = make(chan int, 2)
	fmt.Println(buffChan)
	buffChan <- 1
	buffChan <- 2
	fmt.Println(<-buffChan)
	res, ok := <-buffChan
	fmt.Println(res, ok)

	// 再读取报错了 fatal error: all goroutines are asleep - deadlock!
	// res, ok = <-buffChan
	// fmt.Println(res, ok)

}
