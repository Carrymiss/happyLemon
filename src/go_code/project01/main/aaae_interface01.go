package main

import "fmt"

func main() {
	// 接口
	test87()
	// 多态
	test88()
	// 类型断言
	test89()
}

type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	println("手机开始工作")
}

func (p Phone) Stop() {
	println("手机停止工作")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	println("相机开始工作")
}

func (c Camera) Stop() {
	println("相机停止工作")
}

type Computer struct {
}

// Working 只要实现了usb接口（所谓实现usb接口，就是指实现了usb接口声明所有方法）
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func test87() {
	var c Computer
	// 这里体现了多态的特性
	c.Working(Phone{"vivo"})
	c.Working(Camera{"尼康"})
	// 分割线
	println("-------------分割线--------------")
}

func test88() {
	// 多态的体现 ---》 接口数组
	// 定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}
	fmt.Println(usbArr)
	// 分割线
	println("-------------分割线--------------")
}

type Point01 struct {
	x int
	y int
}

// 类型断言(带检测)
func test89() {
	// a是空接口类型，可以接收任意类型
	var a interface{}
	var point Point01 = Point01{1, 2}
	a = point
	var b Point01
	// 这里的b是Point01类型，要用(&b)
	// 编译器底层做了优化，可以直接用b
	b = a.(Point01)
	fmt.Println(b)
	// 分割线
	println("-------------分割线--------------")
}
