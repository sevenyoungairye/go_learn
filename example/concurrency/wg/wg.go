package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用waitGroup 实现goroutine的同步
// 知道gorutine结束了
var wg sync.WaitGroup

func hello(i int) {
	// 等待gorutine运行完 减一
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func main() {

	for i := 0; i < 5; i++ {
		// 启动一个gorutine就 +1
		wg.Add(1)
		go hello(i)
	}

	wg.Wait() // 等待登记的所有go程结束
}
