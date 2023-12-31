package main

import "fmt"

func main() {
	// 求2数和
	test62()

	// 判断输入的参数是什么类型
	n1 := 1.1
	n2 := 2.2
	n3 := 30
	n4 := "hello"
	n5 := true
	n6 := Calcuator{1.2, 2.3}
	test63(n1, n2, n3, n4, n5, n6, &n6)
}

type Calcuator struct {
	num1 float64
	num2 float64
}

func (cal *Calcuator) test61() float64 {
	// 求和
	return cal.num1 + cal.num2
}

func test62() {
	var cal Calcuator
	cal.num1 = 1.2
	cal.num2 = 2.3
	cal.test61()
	fmt.Printf("两数的和是：%v\n", fmt.Sprintf("%.2f", cal.test61()))
	// 分割线
	println("-------------分割线--------------")
}

func test63(items ...interface{}) {
	// 判断输入的参数是什么类型
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index+1, v)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index+1, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index+1, v)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index+1, v)
		case Calcuator:
			fmt.Printf("第%v个参数是Calcuator类型，值是%v\n", index+1, v)
		case *Calcuator:
			fmt.Printf("第%v个参数是*Calcuator类型，值是%v\n", index+1, v)
		default:
			fmt.Printf("第%v个参数是未知类型，值是%v\n", index+1, v)
		}
	}
	// 分割线
	println("-------------分割线--------------")
}
