package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//通道

func main() {
	//建立连接
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//授权
	if _, err := conn.Do("AUTH", "123456"); err != nil {
		println("false")
		return
	}
	defer conn.Close()
	//通道
	{
		conn.Send("HSET", "student", "name", "wd", "age", "22")
		conn.Send("HSET", "student", "score", "100")
		conn.Send("HGET", "student", "age")
		conn.Flush()
		res1, err1 := conn.Receive()
		fmt.Println(res1, err1)
		res2, err2 := conn.Receive()
		fmt.Println(res2, err2)
		res3, err3 := redis.String(conn.Receive())
		fmt.Printf("%v,%T,%v\n", res3, res3, err3)
	}
}
