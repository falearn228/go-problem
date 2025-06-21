package main

import (
	"fmt"
)

func callCallbacks(a, b func()) {
	go a()
	go b()
}

func main() {
	example()
}

func example() {
	firstDone := make(chan struct{})
	secondDone := make(chan struct{})

	callCallbacks(
		func() {
			fmt.Printf("a")
			close(firstDone)
		},
		func() {
			fmt.Printf("b")
			close(secondDone)
		},
	)

	count := 0
	for count < 2 {
		select {
		case <-firstDone:
			count++
		case <-secondDone:
			count++
		}
	}

	fmt.Printf("%d", count)
}

// 1. a b 2
// 2. b 2
// 3. a 2
// 4. b a 2
// 5. a 2 b
// 6. b 2 a
