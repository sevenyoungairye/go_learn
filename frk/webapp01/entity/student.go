package entity

// https://www.cnblogs.com/remixnameless/p/14318145.html

type Student struct {
	// Id: 编号
	Id int `json:"id" query:"id"`
	// Name: 姓名
	Name string `json:"name" query:"name"`
	// age: 年龄
	Age int `json:"age" query:"age"`
}

func NewStudent(id int, name string, age int) *Student {
	return &Student{Id: id, Name: name, Age: age}
}
