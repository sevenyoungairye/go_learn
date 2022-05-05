package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"top.lel/ginapp01/config"
	"top.lel/ginapp01/dao"
	"top.lel/ginapp01/model"
	"top.lel/ginapp01/tool"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/add", postUser)

	users.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "users")
	})

	users.GET("/comments", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "comments")
	})

	users.GET("/picture", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "picture")
	})

	// https://github.com/gin-gonic/gin#parameters-in-path
	users.DELETE("/del/:id", delUser)
	users.PUT("/updUser", updateUser)
	users.POST("/addUser", saveUser)
	users.GET("/list", getUserList)
}

func delUser(ctx *gin.Context) {
	param := ctx.Param("id")
	fmt.Println("get param...", param)
	id, err := strconv.Atoi(param)
	if err == nil {
		dao.DeleteById(uint(id))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("西内, %s", err))
}

func updateUser(ctx *gin.Context) {
	var user model.UserInfo
	// 绑定到指针
	err := ctx.ShouldBind(&user)
	if err == nil {
		fmt.Println("获取到了用户...", user)
		dao.UpdateUser(user)
		ctx.JSON(http.StatusOK, config.NewBaseResp(http.StatusOK, "ok", ""))
		return
	}
	ctx.JSON(http.StatusInternalServerError, config.NewBaseResp(500, fmt.Sprintf("%s", err), "sth error..."))
}

func saveUser(ctx *gin.Context) {
	user := model.UserInfo{}
	err := ctx.ShouldBind(&user)

	dao.AddUser(user)
	fmt.Println("after do add user... ", user)

	if err == nil {
		resp := config.NewBaseResp(http.StatusOK, "ok", "")
		ctx.JSON(http.StatusOK, resp)
	}
}

func getUserList(ctx *gin.Context) {
	userList := dao.QueryUserList()
	resp := config.NewBaseResp(http.StatusOK, "ok", userList)
	ctx.JSON(http.StatusOK, resp)
}

func postUser(ctx *gin.Context) {
	user := User{}
	err := ctx.ShouldBind(&user)

	fmt.Println("get user...", user)
	fmt.Println(user.CreateTime.Value())
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"user":    user,
		})
	} else {
		fmt.Println("error info...", err)
	}
}

type User struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
	// 2017-01-08T00:00:00Z 默认转换格式...
	Birthday   time.Time `json:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime tool.Date `json:"createTime"`
}
