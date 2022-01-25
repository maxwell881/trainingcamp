package main
//单飞的示例使用
import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"time"
)
var count int32
//下面我们启动 1000 个 Goroutine 去并发调用这两个方法
func main() {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count,-count)
	})
	var(
		wg sync.WaitGroup
		now = time.Now()
		n = 1000
		sg = &singleflight.Group{}
	)
	for i := 0;i<n;i++{
		wg.Add(1)
		go func() {
			//res,_ :=getArticle(1)
			res, _ := singleflightGetArticle(sg, 1)
			if res != "article:1"{
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
}
func singleflightGetArticle(sg *singleflight.Group,id int) (string,error){
	v,err,_ := sg.Do(fmt.Sprintf("%d",id),func()(interface{},error){
		return getArticle(id)
	})
	return v.(string),err
}

func getArticle(id int) (article string,err error){
	atomic.AddInt32(&count,1)
	time.Sleep(time.Duration(count)*time.Millisecond)
	return fmt.Sprintf("article:%d",id),nil
}