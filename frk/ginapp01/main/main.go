package main

import (
	"fmt"
	"top.lel/ginapp01"
)

func main() {

	ginapp01.StartServer()

	func() {
		defer func() {
			fmt.Println("随后就到")
			fmt.Println("结束")
		}()
		fmt.Println("我先执行")
	}()

	//runPerson()
}

func runPerson() {
	p := Person{Name: "jack", Age: 19}
	fmt.Println(p)

	fmt.Println("调用change方法")
	ChangePer(&p)

	fmt.Println("main方法=====")
	fmt.Println(p)

	p.ChangePer(&p)
}

type Person struct {
	Age  int
	Name string
}

// & 是取地址符号 , 即取得某个变量的地址
//*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值

// ChangePer ...
func ChangePer(p *Person) {
	fmt.Println(p, p.Name, p.Age)
	p.Age = 17
	fmt.Println("chang the age, ", p)
}

// ChangePer 对象person的方法
func (p *Person) ChangePer(person *Person) {
	fmt.Println(p == person)
}
