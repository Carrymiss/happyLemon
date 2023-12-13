package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机生成5个数，并将其反转打印
	test41()

	// 将斐波那契数列的前50个数放到切片中
	test42(50)

	// 定义一个二维数组用来保存三个班，每个班五名同学的成绩，并求出每个班的平均分和所有班级的平均分
	test43()

	// 随机生成10个整数(1-100的范围)保存到数组，并倒序打印以及求平均值，求最大值、最大值的下标、最小值、最小值的下标，并查找里面是否有55
	test44()
}

func test41() {
	// 随机生成5个数，并将其反转打印
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	var arr [5]int
	for index, _ := range arr {
		arr[index] = rand.Intn(10000) + 1
	}
	fmt.Println("此时生成是数组是：", arr)

	// 将生成的数组反转
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("此时生成是数组是：", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test42(num int) {
	// 将斐波那契数列的前50个数放到切片中
	var slice1 []uint64 = make([]uint64, num)
	slice1[0] = 1
	slice1[1] = 1
	for i := 2; i < len(slice1); i++ {
		slice1[i] = slice1[i-1] + slice1[i-2]
	}
	fmt.Println("此时生成的切片是：", slice1)
	// 分割线
	println("-------------分割线--------------")
}

func test43() {
	// 定义一个二维数组用来保存三个班，每个班五名同学的成绩，并求出每个班的平均分和所有班级的平均分
	var arr = [3][5]float64{{18, 29, 73, 45, 78}, {17, 24, 34, 46, 65}, {14, 52, 63, 44, 20}}
	// 每个班的成绩和
	var arr2 [3]float64
	// 每个班级的平均分
	var arr3 [3]float64
	for index1, value1 := range arr {
		var sum float64 = 0
		for _, value2 := range value1 {
			sum += value2
		}
		arr2[index1] = sum
		arr3[index1] = sum / float64(len(value1))
	}
	fmt.Println("每个班的成绩和是：", arr2)
	fmt.Println("每个班的平均分是：", arr3)
	// 分割线
	println("-------------分割线--------------")
}

func test44() {
	// 随机生成10个整数(1-100的范围)保存到数组，并倒序打印以及求平均值，求最大值、最大值的下标、最小值、最小值的下标，并查找里面是否有55
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	// 组装数组
	for index, _ := range arr {
		arr[index] = rand.Intn(100) + 1
	}
	fmt.Println("此时生成是数组是：", arr)
	// 倒序打印
	test45(arr)
	// 求最大值
	// 求平均值
	for index, value = range arr {
		sum += value
	}
}
func test45(arr [10]int) {
	// 倒序打印
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("倒叙打印数组为：", arr)
}
