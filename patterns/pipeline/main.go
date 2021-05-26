package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Create the pipeline
	output := MultiplyFloat(Modulus32(SumRand(Rands(15))))
	for v := range output {
		fmt.Printf("output %v\n", v)
	}
}

func Rands(count int) <-chan int {
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

func Modulus32(in <-chan int) <-chan int {
	out := make(chan int)

	go func(in <-chan int, out chan<- int) {
		defer close(out)

		for {
			data, ok := <-in
			if !ok {
				return
			}

			mod := data % 32

			fmt.Printf("modulus %v %% 32 == %v\n", data, mod)
			out <- mod
		}
	}(in, out)

	return out
}

func SumRand(in <-chan int) <-chan int {
	out := make(chan int)

	go func(in <-chan int, out chan<- int) {
		defer close(out)

		for {
			data, ok := <-in
			if !ok {
				return
			}

			r := rand.Int()
			sum := data + r
			fmt.Printf("sum %v + %v == %v\n", data, r, sum)
			out <- sum
		}
	}(in, out)

	return out
}

func MultiplyFloat(in <-chan int) <-chan float64 {
	out := make(chan float64)

	go func(in <-chan int, out chan<- float64) {
		defer close(out)

		for {
			data, ok := <-in
			if !ok {
				return
			}

			out <- float64(data) * rand.Float64()
		}
	}(in, out)

	return out
}
