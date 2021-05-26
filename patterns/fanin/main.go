package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	out := fanin(ctx, rands(500), rands(500), rands(500), rands(500))

	for i := 0; i < 2000; i++ {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-out:
			if !ok {
				continue
			}

			fmt.Printf("received %v on out channel\n", data)
		}
	}
}

func fanin(ctx context.Context, inputs ...<-chan int) <-chan int {
	out := make(chan int)

	for _, i := range inputs {
		go move(ctx, i, out)
	}

	return out
}

func move(ctx context.Context, in <-chan int, out chan<- int) {
	for {

		select {
		case <-ctx.Done():
			return
		case data, ok := <-in:
			if ok {
				select {
				case <-ctx.Done():
					return
				case out <- data:
				}
			}
		}
	}
}

func rands(count int) <-chan int {
	out := make(chan int)

	go func(out chan<- int) {
		defer close(out)

		for i := 0; i < count; i++ {
			value := rand.Int()
			fmt.Printf("input %v\n", value)

			out <- value
		}

	}(out)

	return out
}
