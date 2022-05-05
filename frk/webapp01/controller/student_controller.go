package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"top.lel/main/dao"
	"top.lel/main/entity"
)

// StuList 获取学生列表
func StuList(c echo.Context) error {
	list := dao.StuList()
	fmt.Println(list)
	err := c.JSON(http.StatusOK, list)
	return err
}

// AddStu 新增学生
func AddStu(c echo.Context) error {

	form := c.Request().PostForm
	fmt.Println("form...", form)

	stu := new(entity.Student)
	err := c.Bind(stu)
	if err != nil {
		return err
	}
	fmt.Println("controller get stu... ", stu)

	// stu entity.Student
	dao.AddStu(stu)

	return c.JSON(http.StatusOK, CommonRsp{Code: "200", Msg: "add ok", Data: "ok"})
}

func ModifyStu(c echo.Context) error {
	stu := new(entity.Student)
	err := c.Bind(stu)
	if err != nil {
		return err
	}
	dao.ModifyStu(stu)
	return c.String(http.StatusOK, "modify ok")
}

func DeleteStu(c echo.Context) error {
	pid := c.QueryParam("id")
	fmt.Println("接收到了参数...", pid)
	if pid == "" {
		serviceError := ServiceError{code: "500", msg: "id不为空"}
		return &serviceError
	}
	id, err := strconv.Atoi(pid)
	if err != nil {
		return err
	}
	dao.RemoveStu(id)
	fmt.Println(id)
	return c.String(http.StatusOK, "del ok")
}

func (e *ServiceError) Error() string {

	return fmt.Sprintf("[code: %s, msg: %s]", e.code, e.msg)
}

// ServiceError 实现了Error接口...
type ServiceError struct {
	msg  string
	code string
}

type CommonRsp struct {
	Msg  string      `json:"msg"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}
