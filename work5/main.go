package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"os"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()


	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./t1w.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 定义接收文件读取的字节数组
	var buf [10240]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	t1wstr := string(content)
	fmt.Println(len(t1wstr))
	fmt.Println(t1wstr)

	t5wstr := appendstring(t1wstr,5)
	fmt.Println(len(t5wstr))
	fmt.Println(t5wstr)
	//
	t10wstr := appendstring(t5wstr,2)
	fmt.Println(len(t10wstr))
	fmt.Println(t10wstr)

	t20wstr := appendstring(t10wstr,2)
	fmt.Println(len(t20wstr))
	fmt.Println(t20wstr)

	//写入到redis
	_, err = c.Do("Set", "a20w", t20wstr)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("Get", "a20w"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println("res:")
	fmt.Println(r)

	_, err = c.Do("expire","a20w",20)
	if err != nil{
		fmt.Println(err)
		return
	}
}
func appendstring(s string,n int) string{
	var tostr string
	c := 1
	for c <=n{
		tostr += s
		c ++
	}
	return tostr
}