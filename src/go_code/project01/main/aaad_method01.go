package main

import "fmt"

func main() {
	// 方法的使用
	test72()
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
}

func (p Person) test73() {
	fmt.Println("test73() name=", p.Name)
}
