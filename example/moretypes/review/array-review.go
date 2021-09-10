package main

import "fmt"

/*
	pls see here:

	https://github.com/datawhalechina/go-talent/blob/master/5.数组、切片.md
*/
func init() {
	// testArrChangeVal()
	// foreachArr()
	// initArr()
}

// https://www.runoob.com/go/go-pointers.html

func testArrChangeVal() {
	/*
		go 数组是值传递.
		arr org:  [0 0 0 0 0]
		change..  [9 0 0 0 0]
		after arr change..  [0 0 0 0 0]
	*/
	var arr = [5]int{}
	fmt.Println("arr org: ", arr, &arr)

	noChangeVal(arr)
	fmt.Println("after arr change.. ", arr, &arr)
}

func noChangeVal(arr [5]int) {
	// 修改了值 但是原来的arr值不变
	arr[0] = 9
	fmt.Println("change.. ", arr, &arr)
}

func foreachArr() {
	var arr = [5]int{1, 3, 5, 7, 11}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println("\n使用range函数遍历.. ")
	// index 可以用 "_" 代替..
	for index, value := range arr {
		fmt.Printf("index: %d, value: %d \t", index, value)
	}
	fmt.Println()
}

func initArr() {
	var arr = [5]int{}

	var arr1 = [5]int{1, 2, 3, 4, 5}

	var arr2 = [5]int{3: 10}

	// [0 0 0 0 0] [1 2 3 4 5] [0 0 0 10 0]
	fmt.Printf("%d %d %d \n", arr, arr1, arr2)
}

func main() {

}
