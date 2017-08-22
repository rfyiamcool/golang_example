package main

import (
    "fmt"
    "net"
    "runtime"
    "os"
)

/*
#include <unistd.h>
*/
import "C"

func main() {
    // 守护进程
    C.daemon(1, 1)
    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Println("Starting the server ...")
    fmt.Println(os.Getpid())
    listener, err := net.Listen("tcp", "localhost:8999")
    if err != nil {
        fmt.Println("Error listening", err.Error())
        return
    }

    for {
        _ , err := listener.Accept()
        if err != nil {
            fmt.Println("error accepting", err.Error())
            return
        }
    }
}
