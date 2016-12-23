package main

import (
	"fmt"
	//	"runtime"
	//	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, i := range nums {
			select {
			case out <- i:
			case <-ctx.Done():
				fmt.Println("gen cancel...")
				return
			}
		}
	}()
	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for c := range in {
			select {
			case out <- c * c:
			case <-ctx.Done():
				fmt.Println("square cancel...")
				return
			}
		}
	}()
	return out
}

func merge(ctx context.Context, ins ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(ins))
	out := make(chan int)
	// ERROR: http://studygolang.com/articles/7994
	// REF:   "for"声明中的迭代变量和闭包
	//	for _, in := range ins {
	//		go func() {
	//			for c := range in {
	//				out <- c
	//			}
	//			wg.Done()
	//		}()
	//	}
	// Solution1: New func Outline
	//	ff := func(in <-chan int) {
	//		for c := range in {
	//			out <- c
	//		}
	//		wg.Done()
	//	}
	//	for _, in := range ins {
	//		go ff(in)
	//	}
	// Solution2: Inline func with parameter
	//	for _, in := range ins {
	//		go func(in <-chan int) {
	//			for c := range in {
	//				out <- c
	//			}
	//			wg.Done()
	//		}(in)
	//	}
	// Solution3: Inline func with parameter copy bak
	for _, in := range ins {
		in_copy := in
		go func() {
			defer wg.Done()
			for c := range in_copy {
				select {
				case out <- c:
				case <-ctx.Done():
					fmt.Println("merge cancel...")
					return
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, "Test", "123456")
	//	defer cancelFunc()

	out_new := gen(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	c1 := square(ctx, out_new)
	c2 := square(ctx, out_new)

	mg := merge(ctx, c1, c2)
	fmt.Println(<-mg)
	fmt.Println(<-mg)
	fmt.Println(<-mg)

	cancelFunc()
	time.Sleep(time.Second * 5)
	for {
		if msg, closed := <-mg; !closed {
			fmt.Println("<-mg has closed!")
			return
		} else {
			fmt.Println(msg)
		}
	}
}
