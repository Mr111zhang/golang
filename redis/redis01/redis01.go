package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//建立连接
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// err0 := conn.Send("auth", "123456")
	// fmt.Println(err0)
	//授权
	if _, err := conn.Do("AUTH", "123456"); err != nil {
		println("false")
		return
	}
	println("ok ")
	defer conn.Close()

	//2 设置name为wd
	{
		_, err = conn.Do("SET", "name", "wd")
		if err != nil {
			fmt.Println("redis set error:", err)
		}
	}
	//3:取得name的值
	{
		name, err := redis.String(conn.Do("GET", "name"))
		if err != nil {
			fmt.Println("redis get error:", err)
		} else {
			fmt.Printf("Got name: %s \n", name)
		}
	}
	//4:设置过期时间
	{
		_, err = conn.Do("expire", "name", 10) //10秒过期
		if err != nil {
			fmt.Println("set expire error: ", err)
			return
		}
	}
	//5:批量设置
	{
		_, err = conn.Do("MSET", "name", "wd", "age", 22)
		if err != nil {
			fmt.Println("redis mset error:", err)
		}
		res, err := redis.Strings(conn.Do("MGET", "name", "age"))
		if err != nil {
			fmt.Println("redis get error:", err)
		} else {
			res_type := reflect.TypeOf(res)
			fmt.Printf("res type : %s \n", res_type)
			fmt.Printf("MGET name: %s \n", res)
			fmt.Println(len(res))
		}
	}
	//6:列表操作
	{
		//添加数据 参数 LPUSH（头插）/RPUSH（尾插）  + 表名 + 数据 。。。
		_, err = conn.Do("LPUSH", "list1", "ele1", "ele2", "ele3")
		if err != nil {
			fmt.Println("redis mset error:", err)
		}
		//读取数据
		res, err := redis.String(conn.Do("LPOP", "list1"))
		if err != nil {
			fmt.Println("redis POP error:", err)
		} else {
			res_type := reflect.TypeOf(res)
			fmt.Printf("res type : %s \n", res_type)
			fmt.Printf("res : %s \n", res)
		}

	}
	//7:hash
	{
		//类似键值的方式存入
		_, err1 := conn.Do("HSET", "student", "name", "wd", "age", "22")
		if err1 != nil {
			fmt.Println("err1", err1)
		}
		//取值时 返回的是空接口 需要使用 redis.Int64 转换
		res, err2 := redis.Int64(conn.Do("HGET", "student", "age"))
		if err2 != nil {
			fmt.Println("err2", err2)
		} else {
			res_type := reflect.TypeOf(res)
			fmt.Printf("res type : %s \n", res_type)
			fmt.Printf("res : %d \n", res)
		}

	}
}
