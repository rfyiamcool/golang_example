package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer2 := time.NewTimer(time.Second * 10)
	go func() {
		for {
			select {
			case <-timer2.C:
				fmt.Println("func 1 Timer2 call success .....")
				fmt.Println("Time cost ", time.Now().Sub(start))
			}
		}
	}()
	go func() {
		for {
			select {
			case <-timer2.C:
				fmt.Println("func 2 Timer call success .....")
				fmt.Println("Time cost ", time.Now().Sub(start))
			}
		}
	}()

	fmt.Println(timer2.Reset(time.Second * 2))
	fmt.Println("modify 2 s")

	fmt.Println(timer2.Reset(time.Second * 3))
	fmt.Println("modify 3 s")

	go func() {
		fmt.Println(timer2.Reset(time.Second * 1))
		fmt.Println("modify 1 s")
	}()

	time.Sleep(time.Second * 5)
	fmt.Println(timer2.Reset(time.Second * 1))
	fmt.Println("modify 1 s")

	time.Sleep(300 * time.Second)
}

