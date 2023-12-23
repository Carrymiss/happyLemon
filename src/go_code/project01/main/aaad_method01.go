package main

import "fmt"

func main() {
	// 方法的使用
	test72()
	test75()
	test80()
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
	p.test77(100)
	var n1 = p.test78(100, 1)
	fmt.Println("两数相加的和是：", n1)
	// 分割线
	println("-------------分割线--------------")
}

func (p1 Person) test76() {
	// 计算从1到1000相加
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("从1到1000相加的结果是：", sum)
	// 分割线
	println("-------------分割线--------------")
}

func (p1 Person) test77(n1 int) {
	// 计算从1到入参相加
	var sum int
	for i := 1; i <= n1; i++ {
		sum += i
	}
	fmt.Printf("从1到%v相加的结果是：%v\n", n1, sum)
	// 分割线
	println("-------------分割线--------------")
}

func (p1 Person) test78(n1 int, n2 int) int {
	// 计算从2数的和并返回
	return n1 + n2
}

func test80() {
	var c1 Circle
	c1.Radius = 4.0
	fmt.Println("面积是：", c1.test79())
	// 分割线
	println("-------------分割线--------------")

	var c2 Circle
	c2.Radius = 4.0
	// 这里的c2是指针类型，所以要用(&c2)
	// 编译器底层做了优化，可以直接用c2
	fmt.Println("面积是：", c2.test81())
	// 分割线
	println("-------------分割线--------------")
}

type Circle struct {
	Radius float64
}

func (c1 Circle) test79() float64 {
	return 3.14 * c1.Radius * c1.Radius
}

func (c2 *Circle) test81() float64 {
	// 这里的c2是指针类型，所以要用(*c2)
	// 编译器底层做了优化，可以直接用c2
	return 3.14 * c2.Radius * c2.Radius
}
