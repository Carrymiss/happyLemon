package main

func main() {
	// 接口
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
