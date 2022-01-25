package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"os"
	"t1/t2"
	"time"
)
//声明一个全局的rdb变量
var rdb *redis.Client
var RR = 10
func initClient()(err error){

	rdb = redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
		PoolSize: 100, // 连接池大小
	})
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	_,err = rdb.Ping(ctx).Result()
	return err
}
func main() {
	//fmt.Println(11132323)
	e := initClient()
	if e!= nil{
		fmt.Println("初始化失败")
	}
	a := t2.Da{"aadew",23,43,
	}
	a = t2.Da{A1:"21d",A2:232,A3:3.434}
	fmt.Println(a)
	V8example()

	e = http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		ctx := r.Context()
		// 输出到STDOUT展示处理已经开始
		fmt.Fprint(os.Stdout, "processing request\n")
		// 通过select监听多个channel
		select {
		case <-time.After(2 * time.Second):
			// 如果两秒后接受到了一个消息后，意味请求已经处理完成
			// 我们写入"request processed"作为响应
			w.Write([]byte("request processed"))
		case <-ctx.Done():

			// 如果处理完成前取消了，在STDERR中记录请求被取消的消息
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))

}
func V8example(){
	ctx := context.Background()
	if err := initClient();err != nil{
		return
	}
	//set
	err := rdb.Set(ctx,"key1","value1",0).Err()
	if err != nil{
		panic(err)
	}
	//Get
	val,err := rdb.Get(ctx,"key1").Result()
	if err != nil{
		panic(err)
	}
	fmt.Println(val)
}
