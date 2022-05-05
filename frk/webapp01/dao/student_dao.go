package dao

import (
	"container/list"
	"fmt"
	"top.lel/main/entity"
)

var stuList = list.New()

var stuArr = []*entity.Student{nil}

func init() {
	var stu = entity.NewStudent(1, "jack", 1)
	stuList.PushBack(stu)
	stuList.PushBack(entity.NewStudent(2, "rose", 3))
	stuList.PushBack(entity.NewStudent(3, "sam", 4))

	AddStu(entity.NewStudent(1, "jack", 1))
	AddStu(entity.NewStudent(2, "rose", 3))
	AddStu(entity.NewStudent(3, "sam", 4))
}

func AddStu(student *entity.Student) {
	if student != nil {
		stuList.PushBack(student)
	}

	var addFlg = false
	for i := range stuArr {
		if stuArr[i] == nil {
			stuArr[i] = student
			addFlg = true
			break
		}
	}
	if !addFlg {
		// 数组扩容
		fmt.Println("切片需要被扩容...")
		stuArr = append(stuArr, student)
	}
}

func RemoveStu(id int) {
	for i := range stuArr {
		stu := stuArr[i]
		if stu != nil && stu.Id == id {
			stuArr[i] = nil
		}
	}
}

func ModifyStu(student *entity.Student) {
	for i := range stuArr {
		stu := stuArr[i]
		if stu != nil && stu.Id == student.Id {
			stuArr[i] = student
		}
	}
}

// SearchById not use
func SearchById(id int) entity.Student {
	for i := range stuArr {
		stu := stuArr[i]
		if stu != nil && stu.Id == id {
			return *stu
		}
	}
	return entity.Student{}
}

func StuList() []entity.Student {

	stuSlice := make([]entity.Student, GetAvaLen())
	for i := range stuArr {
		stu := stuArr[i]
		if stu != nil {
			stuSlice[i] = *stu
			fmt.Println(stuSlice[i])
		}
	}
	return stuSlice
}

func GetAvaLen() int {
	var count = 0
	for i := range stuArr {
		stu := stuArr[i]
		if stu != nil {
			count = count + 1
		}
	}
	return count
}
