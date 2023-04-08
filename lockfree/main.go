package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bruceshao/lockfree/lockfree"
)

var (
	goSize    = 10000
	sizePerGo = 10000
)

func main() {
	// lockfree计时
	t := time.Now()
	// 创建事件处理器
	eh := &longEventHandler[uint64]{}
	// 创建消费端串行处理的Lockfree
	lf := lockfree.NewLockfree[uint64](1024*1024, lockfree.Uint8Array, eh,
		lockfree.NewChanBlockStrategy())
	// 启动Lockfree
	if err := lf.Start(); err != nil {
		panic(err)
	}
	// 获取生产者对象
	producer := lf.Producer()
	var wg sync.WaitGroup
	wg.Add(goSize)
	for i := 0; i < goSize; i++ {
		go func(start int) {
			for j := 0; j < sizePerGo; j++ {
				//写入数据
				err := producer.Write(uint64(start*sizePerGo + j + 1))
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("=====lockfree[", time.Now().Sub(t), "]=====")
	fmt.Println("----- lockfree write complete -----")
	time.Sleep(1 * time.Second)
	// 关闭Lockfree
	lf.Close()
}

type longEventHandler[T uint64] struct {
}

func (h *longEventHandler[T]) OnEvent(v uint64) {
	if v%10000000 == 0 {
		fmt.Println("lockfree [", v, "]")
	}
}
