package main

import "fmt"

func main() {
	// map
	test61()
	test62()
	test63()
}

func test61() {
	// 声明map
	var a map[string]string
	// 在使用map前，需要先make，make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	fmt.Println("a的值是：", a)
	// 分割线
	println("-------------分割线--------------")
}

func test62() {
	// map的第二种使用方式
	var a map[string]string = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	fmt.Println("a的值是：", a)
	// 分割线
	println("-------------分割线--------------")

	// map的第三种使用方式
	var a1 map[string]string = map[string]string{
		"no1": "宋江",
		"no2": "吴用",
	}
	fmt.Println("a1的值是：", a1)
	// 分割线
	println("-------------分割线--------------")

	// 使用类型推导
	a2 := map[string]string{
		"no1": "宋江",
		"no2": "吴用",
	}
	fmt.Println("a2的值是：", a2)
	// 分割线
	println("-------------分割线--------------")
}

func test63() {
	// map的增删改查
	// 增加
	var a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "武松"
	fmt.Println("a的值是：", a)
	// 分割线
	println("-------------分割线--------------")

	// 修改
	a["no3"] = "武大郎"
	fmt.Println("a的值是：", a)
	// 分割线
	println("-------------分割线--------------")

	// 删除
	delete(a, "no3")
	fmt.Println("a的值是：", a)
	// 分割线
	println("-------------分割线--------------")

	// 查找
	val, fin := a["no3"]
	if fin {
		fmt.Println("a的值是：", val)
	} else {
		fmt.Println("没有找到no3")
	}
}
