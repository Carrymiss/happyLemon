package main

import "fmt"

func main() {
	// 方法的使用
	test72()
	test75()
}

type Person struct {
	Name string
	Age  int
}

func test72() {
	// 方法的使用
	// 绑定就等于这个Person结构体有了一个test73()的方法
	var p Person
	p.Name = "tom"
	p.Age = 18
	p.test73()
	// 分割线
	println("-------------分割线--------------")
}

func (p Person) test73() {
	fmt.Println("test73() name=", p.Name)
}

func (p1 Person) test74() {
	fmt.Printf("%v是个好人 \n", p1.Name)
	// 分割线
	println("-------------分割线--------------")
}

func test75() {
	var p Person
	p.Name = "tom"
	p.Age = 18
	p.test74()
	p.test76()
}

func (p1 Person) test76() {
	// 计算从1到1000相加
	var sum int
	for i := 1; i < 1000; i++ {
		sum += i
	}
	fmt.Println("从1到1000相加的结果是：", sum)
	// 分割线
	println("-------------分割线--------------")
}