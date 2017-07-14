package main

import (
    "log"
    "time"
    "github.com/influxdata/influxdb/client/v2"
)

const (
    MyDB = "db"         //数据库名
    username = "guest"       //用户名
    password = "guest"   //密码
)

func main(){
    //链接数据库
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://192.168.210.130:8086",
        Username: username,
        Password: password,
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  MyDB,
        Precision: "s",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    // Create a point and add to batch
    tags := map[string]string{"cpu1": "cpu-total1"}
    fields := map[string]interface{}{
        "idle1":   10.1,
        "system1": 53.3,
        "user1":   46.6,
    }
    pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    bp.AddPoint(pt)

    // Write the batch
    c.Write(bp)
}
