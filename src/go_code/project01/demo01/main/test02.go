package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 如果name没有后缀，就加上后缀，否则就返回原来的名字
	test22()

	// 打印金字塔
	test23(10)

	// 打印倒金字塔
	test24(10)

	// 打印空心金字塔
	test25(10)

	// 99乘法表
	test26(9)

	// 数组转置
	test27()

	// 统计执行时间
	test28()

	// 输出当前月份的天数
	test29(2002, 6)

	// 猜数字
	test30()
	// 分割线
	println("-------------分割线--------------")

	// 输出100以内的所有素数，每行显示5个，并求和
	test31()

	// 三天打鱼两天晒网
	test32(2023, 12, 12)

	// 计算器
	test33(10, 20, 1)

	// 输出小写的a-z以及大写的Z-A
	test34()

	// 输入成绩
	test37()

	// 数组打印
	test38()

	// 输出数组的最大值和下标
	test39()

	// 求数组的和和平均值
	test40()
}

func test21(suffix string) func(string) string {
	return func(name string) string {
		// 如果name没有后缀，就加上后缀，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func test22() {
	f := test21(".jpg")
	name1 := f("test")
	name2 := f("bird.jpg")
	fmt.Println("文件名处理后：", name1)
	fmt.Println("文件名处理后：", name2)
	// 分割线
	println("-------------分割线--------------")
}

func test23(num1 int) {
	// 打印金字塔
	for i := 1; i <= num1; i++ {
		for j := 1; j <= num1-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test24(num1 int) {
	// 打印倒金字塔
	for i := 1; i <= num1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*(num1-i)-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test25(num1 int) {
	// 打印空心金字塔
	// 最后一行隔一个打印一个*
	for i := 1; i <= num1; i++ {
		for j := 1; j <= num1-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || (i == num1 && k%3 == 1) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test26(num int) {
	// 99乘法表
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test27() {
	// 给一个二位数组（3*3）转置
	var arr [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arr2 [3][3]int
	// 打印原数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr2[j][i] = arr[i][j]
		}
	}
	// 打印转置后的数组
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Print(arr2[i][j], "\t")
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}

func test28() {
	// 统计执行时间
	start := time.Now().Unix()
	var str string = ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	end := time.Now().Unix()
	fmt.Println("执行时间为：", end-start, "秒")
	// 分割线
	println("-------------分割线--------------")
}

func test29(year int, month int) {
	// 根据年份和月份，打印当月的天数
	var num int = 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		num = 31
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			num = 29
		} else {
			num = 28
		}
	case 4, 6, 9, 11:
		num = 30
	default:
		fmt.Println("输入有误")
	}
	fmt.Printf("%v年%v月有%v天\n", year, month, num)
	// 分割线
	println("-------------分割线--------------")
}

func test30() {
	// 随机猜数字 随机生成一个1-100的数，让用户猜，有10次机会 如果第一次就猜对了，就提示"你真是个天才" 如果第2-3次猜对了，就提示"你很聪明" 如果第4-9次猜对了，就提示"你猜对了" 如果第10次猜对了，就提示"可算猜对了" 如果10次都没猜对，就提示"你太笨了"
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1 // 生成1-100的随机数
	fmt.Println("作弊程序来了，数字是：", target)
	for i := 1; i <= 10; i++ {
		fmt.Print("输入一个猜的数字: ")
		var guess int
		fmt.Scan(&guess)
		if guess == target {
			switch i {
			case 1:
				fmt.Println("你真是个天才")
			case 2, 3:
				fmt.Println("你很聪明")
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("你猜对了")
			case 10:
				fmt.Println("可算猜对了")
			}
			return
		}
	}
	fmt.Println("你太笨了")
}

func test31() {
	// 输出100以内的所有素数，每行显示5个，并求和
	var count int = 0
	for i := 2; i <= 100; i++ {
		flag := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			fmt.Print(i, "\t")
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	// 分割线
	println("-------------分割线--------------")
}

func test32(year int, month int, day int) {
	// 从1990年1月1日开始执行三天打鱼两天晒网，问在以后的某一天中，是打鱼还是晒网
	// 1.计算从1990年1月1日到现在的总天数
	// 2.计算从1990年1月1日到现在的总天数对5取余数
	// 3.如果余数为1、2、3则为打鱼，如果余数为4、0则为晒网
	targetDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	startDate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.Local)
	days := int(targetDate.Sub(startDate).Hours() / 24)
	// 计算相差天数对5取余
	mod := days % 5
	if mod < 3 {
		fmt.Println("在", targetDate.Format("2006年1月2日"), "是打鱼的日子")
	} else {
		fmt.Println("在", targetDate.Format("2006年1月2日"), "是晒网的日子")
	}
	// 分割线
	println("-------------分割线--------------")
}

func test33(num1 float64, num2 float64, operator int) {
	// 计算器
	var result float64
	var operation string
	switch operator {
	case 1:
		result = num1 + num2
		operation = "+"
	case 2:
		result = num1 - num2
		operation = "-"
	case 3:
		result = num1 * num2
		operation = "*"
	case 4:
		if num2 == 0 {
			fmt.Println("除数不能为0！")
			return
		}
		result = num1 / num2
		operation = "/"
	default:
		fmt.Println("无效的操作符！")
		return
	}
	fmt.Printf("计算结果: %.2f %s %.2f = %.2f\n", num1, operation, num2, result)
	// 分割线
	println("-------------分割线--------------")
}

func test34() {
	// 输出小写的a-z以及大写的Z-A
	fmt.Println("小写的 a-z:")
	fmt.Println(test35('a', 'z'))
	fmt.Println("\n大写的 Z-A:")
	lowerCase := test35('a', 'z')
	upperCase := strings.ToUpper(lowerCase)
	reversedUpperCase := test36(upperCase)
	fmt.Println(reversedUpperCase)
}

func test35(start, end byte) string {
	str := ""
	for i := start; i <= end; i++ {
		str += fmt.Sprintf("%c ", i)
	}
	return str
}

func test36(s string) string {
	// 字符串反转
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func test37() {
	// 输入5个成绩，保存到数组中并输出
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%v个成绩：", i+1)
		fmt.Scan(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Printf("第%v个成绩是%v\n", i+1, score[i])
	}
	// 分割线
	println("-------------分割线--------------")
}

func test38() {
	// 创建一个byte类型的26个元素的数组，分别放置'A'-'Z'。使用for循环访问所有元素并打印出来。提示：字符数据运算'A'+1 -> 'B'
	var arr [26]byte
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
	}
	// 使用range遍历
	for _, value := range arr {
		fmt.Printf("%c ", value)
	}
	fmt.Println()
	// 分割线
	println("-------------分割线--------------")
}

func test39() {
	// 输出数组的最大值和其下标
	var arr = [...]int{1, 3, 5, 7, 9}
	maxVal := arr[0]
	maxIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxIndex = i
		}
	}
	fmt.Printf("数组的最大值是%v，下标是%v\n", maxVal, maxIndex)
	// 分割线
	println("-------------分割线--------------")
}

func test40() {
	// 求数组的和和平均值
	var arr = [...]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("数组的和是%.2f，平均值是%.2f\n", sum, sum/float64(len(arr)))

	// 假如是int数组
	var arr2 = [...]int{1, 2, 3, 4, 5}
	sum2 := 0
	for _, value := range arr2 {
		sum2 += value
	}
	fmt.Printf("数组的和是%d，平均值是%.2f\n", sum2, float64(sum2)/float64(len(arr2)))
	// 分割线
	println("-------------分割线--------------")
}
