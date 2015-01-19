package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan time.Time)

	go func() {
		for {
			time.Sleep(time.Nanosecond * 2)
			c <- time.Now()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Nanosecond * 2)
			c <- time.Now()
		}
	}()

	x := 0
	t := time.Now()
	y := t
	for x < 10000 {
		y = t
		t = <-c
		d := t.Sub(y)
		if d.Nanoseconds() < 0 {
			fmt.Print("-")
			panic("!")
		} else {
			fmt.Print("+")
		}
	}
}
