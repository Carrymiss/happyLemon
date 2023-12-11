package main

import "fmt"

func main() {
	// 数组
	test37()
}

func test37() {
	// 定义一个数组
	var hens [6]float64
	// 给数组的每个元素赋值
	hens[0] = 3.0
	hens[1] = 4.0
	hens[2] = 5.0
	hens[3] = 6.0
	hens[4] = 7.0
	hens[5] = 8.0
	// 遍历数组
	weight := 0.0
	for i := 0; i < len(hens); i++ {
		weight += hens[i]
	}
	avgWeight := fmt.Sprintf("%.2f", weight/float64(len(hens)))
	fmt.Println("总体重是", weight, "平均体重是", avgWeight)
	// 分割线
	println("-------------分割线--------------")
}
