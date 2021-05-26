package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	tchan := NewTime(context.Background())

	for i := 0; i < 10; i++ {
		t, ok := <-tchan
		if !ok {
			return
		}

		fmt.Printf("It's %s and all is well!\n", t.Format("3:04:05PM"))
		time.Sleep(time.Second)
	}
}

func NewTime(ctx context.Context) <-chan time.Time {
	tchan := make(chan time.Time)

	go func(tchan chan<- time.Time) {
		defer close(tchan)

		for {

			select {
			case <-ctx.Done():
				return
			case tchan <- time.Now():
			}
		}
	}(tchan)

	return tchan
}
