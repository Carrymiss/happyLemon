package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	//test07()
	//test08()
	//test09()
	//test10()
	//test11()
	//test12()
	//test13()
	//test14()
	//test15()
	i := test16(10)
	fmt.Println("此时斐波那契数列的值是：", i)
	// 分割线
	fmt.Println("-------------分割线--------------")

	j := test17(10)
	fmt.Println("此时数列的值是：", j)
	// 分割线
	fmt.Println("-------------分割线--------------")

	i2 := test18(1)
	fmt.Println("第一天的桃子数量是：", i2)
	// 分割线
	fmt.Println("-------------分割线--------------")

	num := 20
	fmt.Println("main_num=", num)
	test19(&num)
	fmt.Println("num=", num)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 交换两个变量的值
	n1 := 10
	n2 := 20
	fmt.Println("n1=", n1, "n2=", n2)
	test20(&n1, &n2)
	fmt.Println("n1=", n1, "n2=", n2)
}

func test01() {
	// 假如还有97天放假，问：xx个星期零xx天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Println(week, "个星期零", day, "天")
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test02() {
	// 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var huashi float32 = 134.5
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度 %v 对应的摄氏温度为 %v \n", huashi, sheshi)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test03() {
	// 有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, "b=", b)
	// 分割线
	fmt.Println("-------------分割线--------------")
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
	// 分割线
	fmt.Println("-------------分割线--------------")
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
	// 分割线
	fmt.Println("-------------分割线--------------")
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
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test07() {
	// 根据用户输入显示对应的星期时间，如果星期一显示干煸豆角，如果星期二显示鱼香肉丝，如果星期三显示宫保鸡丁，如果星期四显示酸菜鱼，如果星期五显示水煮肉片，如果星期六显示水煮鱼，如果星期日显示烤鸭
	var week string
	fmt.Println("请输入星期")
	fmt.Scanln(&week)
	switch week {
	case "星期一":
		fmt.Println("干煸豆角")
	case "星期二":
		fmt.Println("鱼香肉丝")
	case "星期三":
		fmt.Println("宫保鸡丁")
	case "星期四":
		fmt.Println("酸菜鱼")
	case "星期五":
		fmt.Println("水煮肉片")
	case "星期六":
		fmt.Println("水煮鱼")
	case "星期日":
		fmt.Println("烤鸭")
	default:
		fmt.Println("输入有误")
	}
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test08() {
	// 打印金字塔
	var totalLevel int = 10
	for i := 1; i <= totalLevel; i++ {
		for j := 1; j <= totalLevel-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test09() {
	// 九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Println()
	}
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test10() {
	// 随机生成1-100的数，直到生成99这个数，看看你一共用了几次
	var count int = 0
	for {
		// 设置纳秒级种子
		rand.Seed(time.Now().UnixNano())
		count++
		fmt.Println("用了", count, "次")
		num := rand.Intn(100) + 1 // 生成0-99的随机数
		if num == 99 {
			fmt.Println("生成了99，退出")
			break
		} else {
			fmt.Println("生成的数是", num)
		}
	}
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test11() {
	// 100以内的数求和，求出当和第一次大于20的当前数
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("当前数是", i)
			break
		}
	}
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test12() {
	// 实现登录验证，有三次机会，如果用户名为张无忌，密码为888，提示登录成功，否则提示还有几次机会
	var name string
	var password string
	var count int = 3
	for i := 1; i <= count; i++ {
		fmt.Println("请输入用户名")
		_, _ = fmt.Scanln(&name) // _表示忽略接收的值
		fmt.Println("请输入密码")
		_, _ = fmt.Scanln(&password) // _表示忽略接收的值
		if name == "张无忌" && password == "888" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("还有", count-i, "次机会")
		}
	}
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test13() {
	// 打印1-100之类的奇数 continue实现
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test14() {
	// 从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序 用for break continue实现
	var positiveCount int = 0
	var negativeCount int = 0
	for {
		var num int
		fmt.Println("请输入一个整数")
		_, _ = fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Println("正数的个数是", positiveCount, "负数的个数是", negativeCount)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test15() {
	// 某人有100000元，每经过一次路口，需要进行一次交费，规则如下：当现金>50000时，每次交费5% 当现金<=50000时，每次交费1000 当现金<=0时，结束交费
	// 只使用if实现
	var money float64 = 100000
	var count int = 0
	for {
		if money > 50000 {
			money *= 0.95
			count++
			continue
		}
		if money <= 50000 && money > 0 {
			money -= 1000
			count++
			continue
		}
		if money <= 0 {
			break
		}
	}
	fmt.Println("一共经过了", count, "次路口")
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test16(num int) int {
	// 使用递归求出斐波那契数列的第n个数	1 1 2 3 5 8 13 21 34 55
	if num <= 1 {
		return num
	} else {
		return test16(num-1) + test16(num-2)
	}
}

func test17(num int) int {
	// 有一个数列：f(0)=1,f(1)=4,f(n+2)=2*f(n+1)+f(n),其中n是大于0的整数，求f(10)的值
	if num == 0 {
		return 1
	} else if num == 1 {
		return 4
	} else {
		return 2*test17(num-1) + test17(num-2)
	}
}

func test18(num int) int {
	// 猴子吃桃子问题：猴子第一天摘下若干个桃子，当即吃了一半，还不瘾，又多吃了一个，第二天早上又将剩下的桃子吃掉一半，又多吃了一个，以后每天早上都吃了前一天剩下的一半多一个，到第10天早上想再吃时，见只剩下一个桃子了，求第一天共摘了多少
	if num > 10 || num < 1 {
		fmt.Println("输入有误")
		return 0
	}
	if num == 10 {
		fmt.Println("第10天有1个桃子")
		return 1
	} else {
		return (test18(num+1) + 1) * 2
	}

}

func test19(num *int) {
	*num = *num + 10
	fmt.Println("test19_num=", *num)
}

func test20(n1 *int, n2 *int) {
	// 交换两个变量的值
	*n1, *n2 = *n2, *n1
}
