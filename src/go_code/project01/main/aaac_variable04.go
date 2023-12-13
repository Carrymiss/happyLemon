package main

import "fmt"

func main() {
	// map
	test61()
}

func test61() {
	// 声明map
	var a map[string]string
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	fmt.Print("a的值是：", a)
	// 分割线
	println("-------------分割线--------------")
}
