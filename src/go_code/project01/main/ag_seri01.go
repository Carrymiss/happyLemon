package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 将结构体 序列化
	test97()
	// 将map序列化
	test98()
	// 将切片序列化
	test99()

	// 对基础数据类型序列化
	test100()
}

type Monster1 struct {
	// 指定序列化后的key
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Skill    string  `json:"skill"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
}

func test97() {
	// 将结构体 序列化
	monster := Monster1{
		Name:     "牛魔王",
		Age:      500,
		Skill:    "牛魔拳",
		Birthday: "2011-11-11",
		Sal:      8000.0,
	}

	// 将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
	}
	// 输出序列化后的结果
	fmt.Println("monster序列化后=", string(data))
	// 分割线
	println("-------------分割线--------------")
}

func test98() {
	// 将map序列化
	var a map[string]interface{}
	// 使用map，需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	// 将a序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
	}
	// 输出序列化后的结果
	fmt.Println("a序列化后=", string(data))
	// 分割线
	println("-------------分割线--------------")
}

func test99() {
	// 将切片序列化
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	// 使用map，需要make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	// 使用map，需要make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	slice = append(slice, m2)

	// 将slice序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败 err=", err)
	}
	// 输出序列化后的结果
	fmt.Println("slice序列化后=", string(data))
	// 分割线
	println("-------------分割线--------------")
}

func test100() {
	// 对基础数据类型序列化
	var num1 int = 100
	var num2 float64 = 3.125
	data1, err1 := json.Marshal(num1)
	data2, err2 := json.Marshal(num2)
	if err1 != nil || err2 != nil {
		fmt.Println("序列化失败 err1=", err1)
		fmt.Println("序列化失败 err2=", err2)
	}
	// 输出序列化后的结果
	fmt.Println("num1序列化后=", string(data1))
	fmt.Println("num2序列化后=", string(data2))
	// 分割线
	println("-------------分割线--------------")
}
