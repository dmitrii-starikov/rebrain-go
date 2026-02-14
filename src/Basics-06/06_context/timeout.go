package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	tick := time.NewTicker(time.Millisecond * 50)

	for {
		select {
		case t := <-tick.C:
			fmt.Println(t)
			cancel() // force closes channel
		case <-ctx.Done():
			// chan Done will be closed in 300 ms, so - struct{}{},false - returns
			fmt.Println("context deadline exceeded")
			return
		default:
			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 20)
		}
	}
}
