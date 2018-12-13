package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name   string
		Age    int
		Gender bool
	}

	//1:将Json序列化进入结构体
	var p Person
	var str = "{\"Name\":\"zhuyuqiang\",\"Age\":20,\"Gender\":true}"
	json.Unmarshal([]byte(str), &p)   //
	fmt.Println(p.Age)
	fmt.Println(p.Name)

	//2:将Json序列化到结构体slice
	var ps []Person
	var aJson = "[{\"Name\":\"zhangsan\",\"Age\":25,\"Gender\":true},\n{\"Name\":\"lisi\",\"Age\":20,\"Gender\":false}]"
	json.Unmarshal([]byte(aJson), &ps)
	//[{zhangsan 25 true} {lisi 20 false}] len is 2
	fmt.Println(ps, "len is", len(ps))

	//3:将Json序列化进Map
	var obj map[string]interface{}
	json.Unmarshal([]byte(str), &obj)
	fmt.Println(obj["Name"])
	fmt.Println(obj["Gender"])

	//4:将Json序列化进Slice
	var strs string = `["Go", "Java", "C", "Php"]`
	var aStr []string
	json.Unmarshal([]byte(strs), &aStr)
	//result --> [Go Java C Php]  len is 4
	fmt.Println(aStr, " len is", len(aStr))

	// /*
		//unmarshal to slice
		var strs string = `["Go", "Java", "C", "Php"]`
		var aStr []string
		json.Unmarshal([]byte(strs), &aStr)
		//result --> [Go Java C Php]  len is 4
		fmt.Println(aStr, " len is", len(aStr))
		fmt.Println(aStr[3])
	// */
}
