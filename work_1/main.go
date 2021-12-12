package main

import (
	"database/sql"
	"fmt"
	"os"
	"work_1/dao"
	"work_1/setting"

	ee "github.com/pkg/errors"
)

type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {

	// 先运行  todo.sql 建表
	if len(os.Args) < 2 {
		fmt.Println("config.ini 参数需输入终端")
		return
	}
	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 连接数据库

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接

	
	// 查询数据
	// 模拟model层
	var uu todo

	fmt.Println("=================查询一条数据=================")

	// err2 := dao.DB.QueryRow(`
	// SELECT id,title,status FROM todos WHERE id = ?
	// `, 2).Scan(
	// 	&uu.ID, &uu.Title, &uu.Status,
	// )

	err2 := dao.DB.QueryRow(`
	SELECT id,title,status FROM todos WHERE id = ?
	`, 20).Scan(
		&uu.ID, &uu.Title, &uu.Status,
	)

	switch {
	case err2 == sql.ErrNoRows:
		fmt.Println("models::sql.ErrNoRows", ee.Wrap(err2, "select failed"))
		
	default:
		fmt.Println(uu.ID, uu.Title, uu.Status)
		//业务处理
	}

}
