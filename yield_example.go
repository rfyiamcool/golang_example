package main

import (
	"fmt"
)

func main() {
	c := yield()
	for val := range c {
		fmt.Println("out: %s", val)
		//c <- ""
	}
}

func yield() chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		test := []string{"hello", "world", "1", "2", "3", "4"}
		for _, s := range test {
			c <- s
			//<-c
			fmt.Println("in: %s", s)
		}
	}()

	return c
}
