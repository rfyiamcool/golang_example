package main

import (
    "bytes"
    "crypto/md5"
    "fmt"
    "io"
)

func MakePasswordMd5(s string) string {
    salt1 := "@#$%"
    salt2 := "^&*()"
    buf := bytes.NewBufferString("")

    io.WriteString(buf, salt1)
    io.WriteString(buf, s)
    io.WriteString(buf, salt2)

    t := md5.New()
    io.WriteString(t, buf.String())
    // 输出
    return fmt.Sprintf("%x", t.Sum(nil))
}

func main(){
	fmt.Println(MakePasswordMd5("nb"))
	fmt.Println(MakePasswordMd5("nb"))
	fmt.Println(MakePasswordMd5("nb"))
	fmt.Println(MakePasswordMd5("nb"))
}
