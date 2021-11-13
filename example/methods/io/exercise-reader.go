package main

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// 实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。
// 1. 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func init() {
	fmt.Println("==== exercise.go ====")

	// 使用内置函数构造对象
	var r io.Reader = new(MyReader)
	reader.Validate(r)

}

// Read(p []byte) (n int, err error)

func (r MyReader) Read(p []byte) (n int, err error) {

	return 1, errors.New("sth error")
}
