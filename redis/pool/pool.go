package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 使用redis 建立连接池
var Pool redis.Pool

func init() {
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := Pool.Get()
	//授权
	if _, err := conn.Do("AUTH", "123456"); err != nil {
		println("false")
		return
	}
	//
	res, err := conn.Do("HSET", "student", "name", "jack")
	fmt.Println(res, err)
	res1, err := redis.String(conn.Do("HGET", "student", "name"))
	fmt.Printf("res:%s,error:%v\n", res1, err)
}
