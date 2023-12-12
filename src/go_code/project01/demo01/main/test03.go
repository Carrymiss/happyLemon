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
