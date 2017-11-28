package main

import (
    "fmt"
    "unsafe"
)

func str2bytes(s string) []byte {
    ptr := (*[2]uintptr)(unsafe.Pointer(&s))
    btr := [3]uintptr{ptr[0], ptr[1], ptr[1]}
    return *(*[]byte)(unsafe.Pointer(&btr))
}

func bytes2str(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

func main() {
    s := "xiaorui.cc"
    b := str2bytes(s)
    s2 := bytes2str(b)

    fmt.Println(s, string(b), s2)
}
