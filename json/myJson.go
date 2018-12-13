package main

import (
	"encoding/json"
	"fmt"
)

//
//将结构体序列化为json字符串
type Student struct {
	Name string
	Age  int //字段名首字母一定要大写
}

//
//转json
type Person struct {
	Name        string `json:"username"`
	Age         int
	Gender      bool `json:",omitempty"` //逗号后面表示功能性参数  ， omitempty 表示不初始化忽略这一行
	Profile     string
	Omitcontent string `json:"-"`       //转为json忽略这一行
	Count       int    `json:",string"` //转为string类型
}

//
//
type Person1 struct {
	Name   string
	Age    int
	Gender bool
}

func main() {
	{
		//将json转为结构体切片
		var s []Person1
		var str = "[{\"Name\":\"zhangsan\",\"Age\":30},{\"Name\":\"zhangsan\",\"Age\":40},{\"Name\":\"zhangsan\",\"Age\":50}]"
		json.Unmarshal([]byte(str), &s)
		fmt.Println(s)
		//将json转为 map
		var str1 = "{\"Name\":\"zhangsan\",\"Age\":30}"
		var m map[string]interface{}
		json.Unmarshal([]byte(str1), &m)
		fmt.Println(m)
		fmt.Printf("%T", m["Age"])  // float64 
		//
	}
	//Unmarshal 将 json 转为类型
	{
		var p Person1
		var str = "{\"Name\":\"zhangsan\",\"Age\":30}"
		json.Unmarshal([]byte(str), &p)
		fmt.Println(p.Name)
		fmt.Println(p.Age)
	}
	//marshal 转为 json
	/*
		//struct 转json 的综合运用
		{
			var p *Person = &Person{
				Name:        "zhanyida",
				Age:         20,
				Gender:      true,
				Profile:     "I'm zhangyida",
				Omitcontent: "Omitcontent",
			}
			if bs, err := json.Marshal(p); err != nil {
				panic(err)
			} else {
				fmt.Println(string(bs))
			}
		}
		//将结构体和结构体切片序列化为json字符串
		{
			// s := Student{"zhangsan", 30}    //结构体类型

			s1 := Student{"zhangsan", 30}
			s2 := Student{"zhangsan", 40}
			s3 := Student{"zhangsan", 50}
			s := []Student{s1, s2, s3}                  //切片类型
			if bs, err := json.Marshal(s); err != nil { //Marshal是将任意类型转为 []byte
				panic(err)
			} else {
				fmt.Println(string(bs))
			}
		}
		//将map序列化为json
		{
			m := make(map[string]string)
			m["Go"] = "No1"
			m["PHP"] = "No2"
			if bs, err := json.Marshal(m); err != nil { //Marshal是将任意类型转为 []byte
				panic(err)
			} else {
				fmt.Println(string(bs))
			}

		}
	*/

}
