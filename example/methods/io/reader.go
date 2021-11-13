package main

import (
	"fmt"
	"io"
	"strings"
)

func init() {

	fmt.Println("===== reader.go =====")

	// 返回一个reader
	r := strings.NewReader("hello world..")

	// 准备字节数组 读取r
	b := make([]byte, 8)

	// 将数据读取到byte数组

	for {
		// 每次读到4个字节.. 读到切片b中
		// 返回n读到数组中的字节数, 读到数据流结尾返回EOF错误
		n, err := r.Read(b)
		fmt.Println(n, err, b)

		fmt.Printf("b[:n] : %q \n", b[:n])

		if err == io.EOF {
			break
		}
	}

}

func main() {

}
