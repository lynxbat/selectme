package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		i := 0
		for {
			i++
			go func() { c <- i }()
		}
	}()

	last := 0
	for {
		tt := <-c
		if last < tt {
			fmt.Println(tt)
			return
		}
		last = tt
	}
}
