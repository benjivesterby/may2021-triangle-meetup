package main

import (
	"fmt"
	"math/rand"
)

func main() {
	seven, five, three, other := fanny(rands(50))
	go divisible(7, seven)
	go divisible(5, five)
	go divisible(3, three)
	divisible(-1, other)
}

func fanny(in <-chan int) (<-chan int, <-chan int, <-chan int, <-chan int) {
	seven, five, three, other := make(chan int), make(chan int), make(chan int), make(chan int)

	go func(in <-chan int, seven chan<- int, five chan<- int, three chan<- int, other chan<- int) {
		defer close(seven)
		defer close(five)
		defer close(three)
		defer close(other)

		for {
			value, ok := <-in
			if !ok {
				return
			}

			// Seven divisible
			if value%7 == 0 {
				seven <- value
			} else if value%5 == 0 {
				// Five divisible
				five <- value
			} else if value%3 == 0 {
				// Three divisible
				three <- value
			} else {
				other <- value
			}
		}
	}(in, seven, five, three, other)

	return seven, five, three, other
}

func divisible(divisibility int, in <-chan int) {
	for {
		value, ok := <-in
		if !ok {
			return
		}

		if divisibility < 0 {
			fmt.Printf("%v has unknown divisibility\n", value)
		} else {
			fmt.Printf("%v is divisible by %v\n", value, divisibility)
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

func TypeFan(in <-chan interface{}) (<-chan int, <-chan string, <-chan []byte) {
	ints := make(chan int)
	strings := make(chan string)
	bytes := make(chan []byte)

	go func(in <-chan interface{}, ints chan<- int, strings chan<- string, bytes chan<- []byte) {
		defer close(ints)
		defer close(strings)
		defer close(bytes)

		for {
			data, ok := <-in
			if !ok {
				return
			}

			switch value := data.(type) {
			case int:
				ints <- value
			case string:
				strings <- value
			case []byte:
				bytes <- value
			default:
				fmt.Printf("%T is an unsupported type", data)
			}
		}
	}(in, ints, strings, bytes)

	return ints, strings, bytes
}
