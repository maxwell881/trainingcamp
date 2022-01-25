package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
type Person struct {
	Age int `form:"age" binding:"required,gt=10"`
	Name string `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_form:"2006-01-02" time_utc:"1"`
}
func main() {
	r := gin.Default()

	r.GET("/das", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBind(&person);err != nil{
			context.String(404,fmt.Sprint(err))
			return
		}
		context.String(200,fmt.Sprintf("%#v",person))
	})
	_ = r.Run(":8080")

}
