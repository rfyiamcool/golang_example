package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	RUNNING = true
	wg      sync.WaitGroup
)

func TestPrint() {
	wg.Add(1)
	for {
		if !RUNNING {
			wg.Done()
			break
		}
		fmt.Println("TestPrint")
		time.Sleep(5 * time.Second)
	}

}

func main() {
	go TestPrint()
	go TestPrint()
	go TestPrint()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	select {
	case s := <-c:
		fmt.Println(s)
		RUNNING = false
	}

	if waitTimeout(&wg, time.Second) {
		fmt.Println("超时退出")
	} else {
		fmt.Println("正常退出")
	}

}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(timeout):
		return true
	}
}

