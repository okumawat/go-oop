package main

import (
	"context"
	"fmt"
	"time"

	"github.com/okumawat/go-oop/demo"
)

func doSomething(ctx context.Context, ch *chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done:", ctx.Err())
			return
		case data := <-*ch:
			fmt.Println("Processing:", data)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	// demo.CompositionDemo()
	// demo.FileDemo()
	//demo.WebsocketDemo()
	// ctx := context.Background()
	// ctx = context.WithValue(ctx, "key1", "value1")

	// ctx, cancel := context.WithCancel(ctx)
	// ch := make(chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	ch <- i
	// }
	// go doSomething(ctx, &ch)
	// time.Sleep(500 * time.Millisecond)
	// fmt.Println("cancelling now")
	// cancel()
	demo.Md5Demo("./demo")

}
