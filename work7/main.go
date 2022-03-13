package main

import (
	"fmt"
	"work7/internal/biz"
	"work7/internal/data"
	"work/internal/web"
)

func main() {
	data.Init("configs/config.ini")
	fmt.Println(data.Conf.MySQLConfig)
	err := data.InitMySQL(data.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer data.Close() // 程序退出关闭数据库连接

	// 模型绑定
	data.DB.AutoMigrate(&biz.Todo{})
	// 注册路由
	r := web.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", data.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
