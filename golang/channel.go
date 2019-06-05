package main

import (
	"context"
	"fmt"
	"time"
)

func calcWithSleep(ctx context.Context, num int, sec int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		time.Sleep(time.Duration(sec) * time.Second)

		select {
		case <-ctx.Done():
		case ch <- num + 1:
		}
	}()

	return ch
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	ch1 := calcWithSleep(ctx, 1, 0)
	ch2 := calcWithSleep(ctx, 2, 4)

	for i := 0; i < 2; i++ {
		select {
		case num1 := <-ch1:
			fmt.Println("num1", num1)
			ch1 = nil
		case num2 := <-ch2:
			fmt.Println("num2", num2)
			ch2 = nil
		case <-ctx.Done():
			fmt.Println("timeout")
		}
	}
}
