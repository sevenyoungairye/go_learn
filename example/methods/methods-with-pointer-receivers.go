package main

/*

	1. 创建某个类型的方法时 参数使用值或者指针 但并不应该二者混用
	2. 使用指针作为参数接收者
		首先，方法能够修改其接收者指向的值。

		其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
*/
func init() {

}