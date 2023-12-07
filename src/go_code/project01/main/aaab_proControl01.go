package main

import (
	"fmt"
)

func main() {
	// test03()

	// 循环控制
	//test04()

	// 字符串遍历
	//test05()

	// break用法
	test06()
}
func test03() {
	// switch分支
	var num int
	fmt.Println("请输入一个数字")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("你输入的数字是1")
	case 2:
		fmt.Println("你输入的数字是2")
	case 3:
		fmt.Println("你输入的数字是3")
	default:
		fmt.Println("你输入的数字是其他数字")
	}
}
func test04() {
	// for循环
	for i := 1; i <= 10; i++ {
		// 在输出里面加上i，可以看到i的变化
		fmt.Println("hello world", i)
	}
	// for循环的第二种写法
	var i int = 1
	for i <= 10 {
		fmt.Println("hello world", i)
		i++
	}

	// for循环的第三种写法 无限循环 一般配合break使用
	var j int = 1
	for {
		if j <= 10 {
			fmt.Println("hello world", j)
			j++
		} else {
			break
		}
	}

	// for循环的第四种写法
	var num int = 5
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, i*j)
		}
		fmt.Println()
	}
}

func test05() {
	var str1 string = "hello world"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c \n", str1[i])
	}

	// for range循环
	str := "hello world"
	for index, val := range str {
		fmt.Printf("index=%v val=%c \n", index, val)
	}

	// 如果只想获取到字符串的值，不想获取到下标，可以使用_来忽略
	for _, val := range str {
		fmt.Printf("val=%c \n", val)
	}

	// 使用切片遍历带有中文的字符串
	str2 := "hello 北京"
	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("字符=%c \n", str3[i])
	}

	// 使用for range遍历带有中文的字符串
	for index, val := range str3 {
		fmt.Printf("index=%v val=%c \n", index, val)
	}
}

func test06() {
	// break语句
	// break会根据就近原则，默认跳转出最近的循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 2 {
				break
			} else {
				fmt.Println("j=", j)
			}
		}
		fmt.Println("i=", i)
	}
	fmt.Println("-----------分割符---------------")
	// break语句还可以跳出指定的循环
label1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 2 {
				break label1
			} else {
				fmt.Println("j=", j)
			}
		}
		fmt.Println("i=", i)
	}
	fmt.Println("-----------分割符---------------")
	for i := 0; i < 10; i++ {
	label2:
		for j := 0; j < 10; j++ {
			if j > 2 {
				break label2
			} else {
				fmt.Println("j=", j)
			}
		}
		fmt.Println("i=", i)
	}
}
