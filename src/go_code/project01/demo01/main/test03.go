package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机生成5个数，并将其反转打印
	test41()
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
