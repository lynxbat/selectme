package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan time.Time)
	c2 := make(chan time.Time)

	go func() {
		for {
			time.Sleep(time.Nanosecond * 2)
			c1 <- time.Now()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Nanosecond * 2)
			c2 <- time.Now()
		}
	}()

	x := 0
	t := time.Now()
	y := t
	for x < 10000 {
		select {
		case tt := <-c1:
			y = t
			t = tt
			d := t.Sub(y)
			if d.Nanoseconds() < 0 {
				fmt.Print("-")
				panic("!")
			} else {
				fmt.Print("+")
			}
		case tt := <-c2:
			y = t
			t = tt
			d := t.Sub(y)
			if d.Nanoseconds() < 0 {
				fmt.Print("-")
				panic("!")
			} else {
				fmt.Print("+")
			}
		}

	}
}
