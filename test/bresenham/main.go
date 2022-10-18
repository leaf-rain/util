package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var ctx = context.WithValue(context.Background(), 1, 1)
	go func(ctx context.Context) {
		for {
			fmt.Println(ctx.Value(1))
			time.Sleep(time.Second)
		}
	}(ctx)
	time.Sleep(time.Second * 3)

}
