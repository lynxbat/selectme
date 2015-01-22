package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	i := 0
	go func() {
		for {
			i++
			if i%2 == 0 {
				go func() { c2 <- i }()
			} else {
				go func() { c1 <- i }()
			}
		}
	}()

	last := 0
	for {
		select {
		case tt := <-c1:
			if tt < last {
				fmt.Println(tt)
				return
			}
			last = tt
		case tt := <-c2:
			if tt < last {
				fmt.Println(tt)
				return
			}
			last = tt
		}
	}
}
