package main

import (
    "fmt"

    "github.com/golang/snappy"
)

const test = `{"Ag(T+D)":{"instID":"Ag(T+D)","name":"白银延期","last":"4141","upDown":"21","upDownRate":"0.51","quoteDate":"20170328","quoteTime":"22:34:29"},"Au(T+D)":{"instID":"Au(T+D)","name":"黄金延期","last":"280.55","upDown":"0.88","upDownRate":"0.31","quoteDate":"20170328","quoteTime":"22:34:15"},"mAu(T+D)":{"instID":"mAu(T+D)","name":"Mini黄金延期","last":"280.5","upDown":"0.7","upDownRate":"0.25","quoteDate":"20170328","quoteTime":"22:34:10"}}`

func main() {

    fmt.Println("source len:", len(test))

    got := snappy.Encode(nil, []byte(test))
    fmt.Println("compressed len:", len(got))

    a, _ := snappy.Decode(nil, got)
    fmt.Println("uncompressed len:", len(a))

}
