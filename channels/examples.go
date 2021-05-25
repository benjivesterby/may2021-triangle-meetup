package channels

import (
	"context"
	"time"
)

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
