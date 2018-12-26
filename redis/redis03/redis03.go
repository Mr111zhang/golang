package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

//观察者
func main() {
	go Subs()
	time.Sleep(time.Second * 1)
	go Push("this is wd")
	time.Sleep(time.Second * 1)
	go Push("this is wd2")
	time.Sleep(time.Second * 1)
}

func Subs() {
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

	//建立订阅者  用于监听 发布者
	psc := redis.PubSubConn{conn}
	psc.Subscribe("channel1") //订阅channel1频道
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s:message:%s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s:%s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func Push(message string) {
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

	_, err1 := conn.Do("PUBLISH", "channel1", message)
	if err1 != nil {
		fmt.Println("pub err", err1)
		return
	}
	defer conn.Close()
}
