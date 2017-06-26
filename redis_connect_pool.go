package main

import (
//	"fmt"
	"github.com/garyburd/redigo/redis"
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)
var MAX_POOL_SIZE = 20

var redisPool chan redis.Conn

func putRedis(conn redis.Conn) {
	if redisPool == nil {
		redisPool = make(chan redis.Conn, MAX_POOL_SIZE)
	}
	if len(redisPool) >= MAX_POOL_SIZE {
		conn.Close()
		return
	}
	redisPool <- conn
}

func InitRedis(network, address string) redis.Conn {
	redisPool = make(chan redis.Conn, MAX_POOL_SIZE)
	if len(redisPool) == 0 {
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				c, err := redis.Dial(network, address)
				if err != nil {
					panic(err)
				}
				putRedis(c)
			}
		}()
	}
	return <-redisPool
}


func main() {
	fmt.Println()

	c := InitRedis("tcp", "127.0.0.1:6379")

	//test uuid
	fmt.Println(time.Now())
	startTime := time.Now()

	var Success, Failure int
	for i := 0; i < 100000; i++ {
		if ok, _ := redis.Bool(c.Do("HSET", "payVerify:session", uuid.New(), "aaaa")); ok {
			Success++
			// break
		} else {
			Failure++
		}
	}
	fmt.Println(time.Now())
	fmt.Println("用时：", time.Now().Sub(startTime), "总计：100000,成功：", Success, "失败：", Failure)
}
