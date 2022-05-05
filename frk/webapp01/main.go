package main

import (
	"net/http"
	"top.lel/main/controller"
)
import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/**
https://github.com/labstack/echo
*/

func main() {
	httpServer()
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func httpServer() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	e.POST("/stu/add", controller.AddStu)
	e.GET("/stu/list", controller.StuList)
	e.DELETE("/stu/del", controller.DeleteStu)
	e.PUT("/stu/upd", controller.ModifyStu)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type User struct {
	// id: 编号
	id int
	// name: 姓名
	name string
	// age: 年龄
	age int
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func NewUser(id int, name string, age int) *User {
	return &User{id: id, name: name, age: age}
}
