// 通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。

// 例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。

package main

import (
	"fmt"
)

// 声明类型IPAddr 长度为4的byte数组
type IPAddr [4]byte

// IPAddr 添加一个 "String() string" 方法
// 重写了type IPAddr的toString方法
func (addr IPAddr) String() string {

	res := ""

	for i, v := range addr {
		tmp := fmt.Sprintf("%d", v)

		if i == len(addr)-1 {
			res = res + tmp
		} else {
			res = res + tmp + "."
		}

	}

	return res
}

func init() {

	fmt.Println("===== exercise-Stringer.go =====")

	// var m map[string]string = map[string]string{}

	m := map[string]IPAddr{
		"localhost": {127, 0, 0, 1},
		"googleDns": {8, 8, 8, 8},
	}

	for key, value := range m {
		fmt.Printf("key: %v, value: %v \n", key, value)
	}

	tmp := IPAddr{8, 8, 8, 8}
	fmt.Println(tmp)

	/*

		// byte的范围..
		var b [4]byte = [4]byte{'^', 0, 255}
		fmt.Println(b, len(b), cap(b))

		var ip IPAddr = IPAddr{8, 8, 8, 8}
		fmt.Println(ip)

	*/
}

func main() {

}
