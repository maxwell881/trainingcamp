package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("got the stop channel2")
				return
			default:
				fmt.Println("still working2")
				time.Sleep(1*time.Second)
			}
		}
	}()
	time.Sleep(5*time.Second)
	fmt.Println("stop the gorutine2")
	cancel()
	time.Sleep(2*time.Second)
}
