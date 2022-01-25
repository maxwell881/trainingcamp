package main
//单飞的示例使用
import (
	"database/sql"
	"fmt"
	"golang.org/x/sync/singleflight"
	"strings"
	"sync"
	"time"

	//"strings"
	//"sync/atomic"
	_ "github.com/go-sql-driver/mysql"

)
//var counts int32
// 定义一个全局对象db
var db *sql.DB
type user struct {
	id   int
	currentid  int64
	ti string
}
//下面我们启动 1000 个 Goroutine 去并发调用这两个方法
func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("链接成功！")
	var(
		wg sync.WaitGroup
		now = time.Now()
		n = 1000
		sg = &singleflight.Group{}
	)
	for i := 0;i<n;i++{
		wg.Add(1)
		go func() {
			//str,_ := getdata(13181)
			str,_ := singleflightGetData(sg,13181)

			fmt.Println(str)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))

}
func singleflightGetData(sg *singleflight.Group,id int) (string,error){
	v,err,_ := sg.Do(fmt.Sprintf("%d",id),func()(interface{},error){
		return getdata(id)
	})
	return v.(string),err
}

func getdata(id int) (u string,err error) {
	sqlStr := "select id, currentId, runningTime from savedId where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.currentid, &u.ti)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("id:%d ti:%s currentid:%d\n", u.id, u.ti, u.currentid),nil
	}
	return 
}



func initDB() (err error) {
	// DSN:Data Source Name
	//dsn := "root:Dachangtui123$%^@tcp(192.168.19.92:3306)/ai_test?charset=utf8&parseTime=True"

	//或者使用
	userName := "root"
	password := "Dachangtui123$%^"
	ip := "192.168.19.92"
	port := "3306"
	dbName := "ai_test"
	dsn := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	// fmt.Println("dsn:","连接数据succeed")

	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}