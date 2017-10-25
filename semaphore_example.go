package main
import(
    "fmt"
    "time"

    nsema "github.com/niean/gotools/concurrent/semaphore"
)

func main(){
    // init sema
    concurrentNum := 2 // 5,10..
    sema := nsema.NewSemaphore(concurrentNum)

    // use sema
    for i:=0; i<100; i++ {
        go func(num int){
            if !sema.TryAcquire() {
                fmt.Printf("%d, get sema, fail\n", num)
                return
            }
            defer sema.Release()

            time.Sleep(1*time.Nanosecond)
            fmt.Printf("%d, get sema, ok\n", num)
        }(i)
    }

    // keep alive
    time.Sleep(2*time.Second)
}
