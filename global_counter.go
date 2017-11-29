package main

import (
    "fmt"

    "sync"
)

type single struct {
    mu     sync.Mutex
    values map[string]int64
}

var counters = single{
    values: make(map[string]int64),
}

func (s *single) Get(key string) int64 {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.values[key]
}

func (s *single) Incr(key string) int64 {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.values[key]++
    return s.values[key]
}

func main() {
    fmt.Println(counters.Incr("bar"))
    fmt.Println(counters.Incr("bar"))
    fmt.Println(counters.Incr("bar"))

    fmt.Println(counters.Get("foo"))
    fmt.Println(counters.Get("bar"))
}
