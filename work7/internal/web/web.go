package web

import (
	"github.com/gin-gonic/gin"
	"work3/internal/biz"
	"work3/internal/data"
)

func SetupRouter() *gin.Engine {
	if data.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", biz.IndexHandler)

	// v1

	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", biz.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", biz.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", biz.UpdateATodo2)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", biz.DeleteATodo2)
	}
	return r
}