package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var errs []error
	chErrs := make(chan error, 3)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	for i := 0; i < 3; i++ {
		go func(timeout int) {
			time.Sleep(time.Second * time.Duration(timeout))
			chErrs <- nil
		}(i)
	}

	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			errs = append(errs, ctx.Err())
		case err := <-chErrs:
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
	}

	fmt.Println("Done!")
}
