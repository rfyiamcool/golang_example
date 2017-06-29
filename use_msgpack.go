package main

import (
    "fmt"
    "time"

    "encoding/json"
    "gopkg.in/vmihailenco/msgpack.v2"
)

type ts struct {
    C   string
    K   string
    T   int
    Max int
    Cn  string
}

func main() {

    var in = &ts{
        C:   "LOCK",
        K:   "31uEbMgunupShBVTewXjtqbBv5MndwfXhb",
        T:   1000,
        Max: 200,
        Cn:  "中文",
    }

    ExampleMsgpack(in)
    ExampleJson(in)
}

func ExampleJson(in *ts) {

    t1 := time.Now()

    for i := 0; i < 100000; i++ {
        // encode
        b, _ := json.Marshal(in)
        // decode
        var out = ts{}
        _ = json.Unmarshal(b, &out)
    }
    t2 := time.Now()
    fmt.Println("Json 消耗时间：", t2.Sub(t1), "秒")
}

func ExampleMsgpack(in *ts) {

    t1 := time.Now()

    for i := 0; i < 100000; i++ {
        // encode
        b, _ := msgpack.Marshal(in)
        // decode
        var out = ts{}
        _ = msgpack.Unmarshal(b, &out)
    }
    t2 := time.Now()
    fmt.Println("msgpack 消耗时间：", t2.Sub(t1), "秒")
}
