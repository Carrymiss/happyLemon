package main

import "fmt"

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}

func test01() {
	// 假如还有97天放假，问：xx个星期零xx天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Println(week, "个星期零", day, "天")
}

func test02() {
	// 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var huashi float32 = 134.5
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度 %v 对应的摄氏温度为 %v \n", huashi, sheshi)
}
func test03() {
	// 有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, "b=", b)
}

func test04() {
	// 求两个数的最大值
	var num1 int = 10
	var num2 int = 20
	var max int
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}
	fmt.Println("max=", max)
}

func test05() {
	// 输入年龄，如果大于18岁，输出你年龄大于18，要对自己的行为负责
	// 如果小于18岁，输出你的年龄小于18，快去学习
	var age int
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你年龄大于18，要对自己的行为负责")
	} else {
		fmt.Println("你的年龄小于18，快去学习")
	}
}

func test06() {
	// 接受一个字符，表示用户的评价，A、B、C、D、E，如果是A，输出评价优秀，如果是B，输出评价良好，如果是C，输出评价及格，如果是D，输出评价不及格，如果是E，输出评价不合格，其他输出评价输入有误
	var score string
	fmt.Println("请输入评价")
	fmt.Scanln(&score)
	switch score {
	case "A":
		fmt.Println("评价优秀")
	case "B":
		fmt.Println("评价良好")
	case "C":
		fmt.Println("评价及格")
	case "D":
		fmt.Println("评价不及格")
	case "E":
		fmt.Println("评价不合格")
	default:
		fmt.Println("评价输入有误")
	}
}
