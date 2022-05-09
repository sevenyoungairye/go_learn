package dao

import (
	"fmt"
	"top.lel/ginapp01/config"
	"top.lel/ginapp01/model"
)

var (
	db, _ = config.GetGlobalDB()
)

func DeleteById(id uint) {
	db.Delete(model.UserInfo{ID: id})
}

func AddUser(u model.UserInfo) {
	//db.Save(u)

	db.Create(&u)
}

func UpdateUser(u model.UserInfo) {
	db.Model(&u).Updates(u)
}

func QueryUserList() []model.UserInfo {
	// 声明切片
	var userList []model.UserInfo
	// preload启用预加载模式, 会把结果带出来.
	res := db.Preload("Company").Find(&userList)
	fmt.Println("查询到结果: ", *res)
	/*res.Scan(&userList)
	fmt.Println("用户列表, ", userList)*/
	for idx := range userList {
		userList[idx].MemberNumber.Valid = false
		userList[idx].MemberNumber.String = ""
	}
	return userList
}
