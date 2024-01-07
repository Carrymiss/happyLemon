package main

import "encoding/json"

func main() {
	// 将json字符串，反序列化成struct
	test101()
	// 将json字符串，反序列化成map
	test102()
	// 将json字符串，反序列化成切片
	test103()
}
func test101() {
	// 将json字符串，反序列化成struct
	str := `{"age":18,"name":"jack","address":"北京"}`
	var p Person1
	// 反序列化
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		println("反序列化失败")
	}
	println(p.Name, p.Age, p.Address)
	// 分割线
	println("-------------分割线--------------")
}

type Person1 struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func test102() {
	// 将json字符串，反序列化成map
	str := `{"age":18,"name":"jack","address":"北京"}`
	// 这里不需要make map，因为Unmarshal函数内部会make
	var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		println("反序列化失败")
	}
	println(m)
	// 分割线
	println("-------------分割线--------------")
}

func test103() {
	// 将json字符串，反序列化成切片
	str := `[{"age":18,"name":"jack","address":"北京"},{"age":19,"name":"tom","address":"上海"}]`
	// 这里不需要make map，因为Unmarshal函数内部会make
	var slice []map[string]interface{}
	// 反序列化
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		println("反序列化失败")
	}
	println(slice)
	// 分割线
	println("-------------分割线--------------")
}
