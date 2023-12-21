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
	test68()
	test69()
	test70()
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
	// 分割线
	println("-------------分割线--------------")
}

func test68() {
	// 声明结构体的方式
	var stu1 = stu{"tom", 18, "北京"}
	fmt.Println("stu1的值是：", stu1)

	// 声明结构体的方式 这种方式声明出来的是结构体的指针类型
	var stu2 = new(stu)
	(*stu2).Name = "mary"
	(*stu2).Age = 28
	(*stu2).Address = "上海"
	fmt.Println("stu2的值是：", *stu2)

	// go语言的作者优化了上面的申明 也可以用如下方式
	// 在底层加上了(*stu3) 也就是取值运算
	var stu3 = new(stu)
	stu3.Name = "mary"
	stu3.Age = 28
	stu3.Address = "上海"
	fmt.Println("stu3的值是：", *stu3)

	// 结构体的声明
	var stu4 = &stu{"mary", 28, "上海"}
	fmt.Println("stu4的值是：", *stu4)
	var stu5 = &stu{}
	stu5.Name = "mary"
	stu5.Age = 28
	stu5.Address = "上海"
	fmt.Println("stu5的值是：", *stu5)

	// 分割线
	println("-------------分割线--------------")
}

func test69() {
	// 结构体字段在内存中是连续存在的
	var stu1 = stu{"tom", 18, "北京"}
	fmt.Printf("stu1的Name的地址是：%p\n", &stu1.Name)
	fmt.Printf("stu1的Age的地址是：%p\n", &stu1.Age)
	fmt.Printf("stu1的Address的地址是：%p\n", &stu1.Address)
	// 分割线
	println("-------------分割线--------------")
}

type Rect struct {
	leftUp, rightDown Point
}

type Point struct {
	x int
	y int
}

func test70() {
	// 结构体嵌套
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println("r1的值是：", r1)
	fmt.Printf("r1的leftUp,x的地址是：%p\n", &r1.leftUp.x)
	fmt.Printf("r1的leftUp,y的地址是：%p\n", &r1.leftUp.y)
	fmt.Printf("r1的rightDown,x的地址是：%p\n", &r1.rightDown.x)
	fmt.Printf("r1的rightDown,y的地址是：%p\n", &r1.rightDown.y)
	// 分割线
	println("-------------分割线--------------")
}
