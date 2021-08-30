package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("hah, %g %s \n", math.Pi, "hello world..")
	fmt.Printf("value.. %g \n", math.Sqrt(9))

	name := "jack"

	var s = greet(name)
	fmt.Println(s)

	fmt.Println("swap sth start...")
	s1, s2 := swap("哈哈", "jack")
	fmt.Println("get swap val...", s1, s2)

	a1, a2 := split(1500)
	fmt.Println("split func..", a1, a2)

	paramDemo()

	initVar()

	basicType()

	typeConvert()

	typeTransfer()

	constDemo()

	numConst()

}

const (
	Big = 1 << 100
	// 1 << 1
	Small = Big >> 99
)

func numConst() {
	fmt.Println(Small)

	fmt.Println(1<<2, 1<<3)
	fmt.Println("====数值常量====")
	fmt.Println(needInt(Small), needFlat(Small), needFlat(Big))
}

func needFlat(f float64) float64 {

	return f * 0.1
}

func needInt(num int) int {

	return num*10 + 1
}

func constDemo() {
	const Hello = "hello"
	fmt.Println(Hello)
}

// 类型传递
func typeTransfer() {
	var i int
	j := i
	k := "hah"
	l := k
	fmt.Printf("type of j, %T %T\n", j, l)
}

func typeConvert() {
	// 转换必须显示转换..
	var i int = 4
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)
}

func basicType() {

	var (
		Name   string = "hah"
		ToBe   bool   = false
		MaxInt uint64 = 1<<64 - 1
	)

	fmt.Println(Name, ToBe, MaxInt)

	/*
		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // uint8 的别名

		rune // int32 的别名
			// 表示一个 Unicode 码点

		float32 float64

		complex64 complex128
	*/

}

func initVar() {
	// var 变量名 类型
	var d float32 = 4.5
	fmt.Println(d)
	// 变量初始化, 可以根据给定的值获取对应的类型
	var a, b = "ha", true
	fmt.Println("init val", a, b)

	// := show a variable.. 必须在函数内使用..
	c := 1
	fmt.Println("the c of value.. ", c)
}

// var 出现在包或者函数级别
var c, python, java bool

func paramDemo() {
	// i, c, python, java... 都未被初始化 但是默认值 他们叫零值
	var i int
	fmt.Printf("hah %b, %d, %t \n", i, i, python)
}

// 可以省略返回值, 用return代替
// 但是可读性降低了
func split(sum int) (c, d int) {
	c = sum * 4 / 9
	d = sum - c
	return
}

// the return value just return the value of type..
// param is x, y..
func swap(x, y string) (string, string) {

	fmt.Printf("current val... %s, %s \n", x, y)

	return y, x
}

func greet(name string) string {

	return "hello " + name + ", what's up."
}
