package main

import "fmt"

func main() {
	// map
	test61()
	test62()
	test63()
	test64()

	// 结构体
	test66()
	test67()
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

func test64() {
	// map是引用类型
	// map 回自动扩容
	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 80
	map1[3] = 70
	map1[4] = 60
	test65(map1)
	fmt.Println("map1的值是：", map1)
	// 分割线
	println("-------------分割线--------------")
}

func test65(map1 map[int]int) {
	map1[1] = 100
}

type stu struct {
	Name    string
	Age     int
	Address string
}

func test66() {
	students := make(map[string]stu, 10)
	// 创建两个学生
	stu1 := stu{"tom", 18, "北京"}
	stu2 := stu{"mary", 28, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println("students的值是：", students)

	for k, v := range students {
		fmt.Printf("学生的编号是：%v，学生的姓名是：%v，学生的年龄是：%v，学生的地址是：%v\n", k, v.Name, v.Age, v.Address)
	}
	// 分割线
	println("-------------分割线--------------")
}

func test67() {
	var stu1 stu
	stu1.Name = "tom"
	stu1.Age = 18
	stu1.Address = "北京"
	stu2 := stu1
	stu2.Name = "mary"
	fmt.Println("stu1的值是：", stu1)
	fmt.Println("stu2的值是：", stu2)
}
