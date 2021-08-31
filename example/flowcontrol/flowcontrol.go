package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	forCal()

	fmt.Printf("ifCondition(99): %v\n", ifCondition(100))
	fmt.Printf("ifCondition(-4): %v\n", ifCondition(-4))

	fmt.Printf("ifSimple(2, 3, 20): %v\n", ifSimple(2, 3, 20))
	fmt.Printf("ifSimple(3, 3, 20): %v\n", ifSimple(3, 3, 20))

	fmt.Printf("ifElseDemo(2, 3, 20): %v\n", ifElseDemo(2, 3, 20))
	fmt.Printf("ifElseDemo(3, 3, 20): %v\n", ifElseDemo(3, 3, 20))

	fmt.Printf("recursion(-9): %v\n", recursion(-9))

	switchDemo()

	switchSeq()

	switchTrue()

	deferDemo()

	deferStack()
}

func deferStack() {

	fmt.Println("counting...")

	for i := 0; i < 10; i++ {
		// 压栈..
		defer fmt.Println("i", i)
	}

	fmt.Println("done..")
}

func deferDemo() {

	// defer 返回值类似压栈操作..
	// defer 不会优先执行， 但不影响recursion方法的执行顺序
	defer fmt.Println(recursion(-9))

	defer fmt.Println("world")

	fmt.Println("hello ")
}

func switchTrue() {
	// if then else..
	// switch true, 没得条件, 条件在case里面
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good moorning..")
	case t.Hour() < 17:
		fmt.Println("Good afternoon..")
	default:
		fmt.Println("Good evening")
	}
}

func switchSeq() {
	fmt.Println("What is Saturday..?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		{
			fmt.Println("Today !")
		}
	case today + 1:
		{
			fmt.Println("Tomorrow.")
		}
	case today + 2:
		{
			fmt.Println("In two days.")
		}
	default:
		fmt.Println("Too far away.")
	}
}

// switch 自带break
func switchDemo() {

	fmt.Println("Go OS start.. ")

	// 可以简写表达式..
	switch os := runtime.GOOS; os {
	case "linux":
		{
			fmt.Println("the os is ", os)
		}
	case "MacOS":
		{
			fmt.Println("the os is ", os)
		}
	default:
		{
			fmt.Println("default.. the os is ", os)
		}
	}
}

func recursion(i int) int {

	if i > 0 {
		return i
	}

	fmt.Print(i, " ")

	i++

	return recursion(i)
}

func ifElseDemo(a, n, lim float64) float64 {

	if v := math.Pow(a, n); v < lim {

		return v
	} else {
		fmt.Printf("g > v... %g >= %g\n", v, lim)
	}

	// v的作用域结束..
	return lim
}

// m³ if 里面也可以像for一样 声明简单的语句
// v的作用域在if块里有用
func ifSimple(a, n, lim float64) float64 {
	if v := math.Pow(a, n); v < lim {

		return v
	}

	return lim
}

func ifCondition(num float64) float64 {
	if num < 0 {
		return math.Sqrt(-num)
	}

	return math.Sqrt(num)
}

func forCal() {
	var sum int = 0
	// 省略小括号
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 初始化语句和后置语句可以省略
	/*
		for ;sum < 10; {
			sum += sum
		}

		sum = 0
		for sum < 10 {
			// to do sth...
			// 不论sum做出何种改变.. 都是死循环状态..
		}
	*/

	// for -> while 循环
	sum = 1
	for sum < 10 {
		sum += sum
		fmt.Println("~~~", sum)
	}
	fmt.Println(sum)

	// 无限循环 for {}
	var i int = 1
	for {
		fmt.Println("康娜酱是小可爱~~ (●ˇ∀ˇ●)")
		i++
		if i > 10 {
			break
		}
	}
}
