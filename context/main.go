package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func main() {
	//	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, "Test", "123456")
	//	defer cancelFunc()

	if t, ok := ctx.Deadline(); ok {
		fmt.Println(time.Now())
		fmt.Println(t.String())
	}
	go func(ctx context.Context) {
		fmt.Println(ctx.Value("Test"))
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
				//			default:
				//				continue
			}
		}
	}(ctx)
	//	if ctx.Err() == nil {
	//		fmt.Println("Sleep 10 seconds...")
	//		time.Sleep(time.Second * 10)
	//	}
	//	if ctx.Err() != nil {
	//		fmt.Println("Alredy exit...")
	//	}
	time.Sleep(time.Second * 3)
	cancelFunc()
	//	for {
	//		if ctx.Err() != nil {
	//			fmt.Println("gracefully exit...")
	//			break
	//		}
	//	}
}
