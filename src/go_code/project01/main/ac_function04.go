package main

import "fmt"

func main() {
	// 冒泡排序
	test54()

	// 二分查找法
	var arr = [5]int{1100, 54, 387, 894, 5798}
	test55(&arr, 0, len(arr)-1, 1100)
}

func test54() {
	// 冒泡排序
	var arr = [5]int{32312, 43, 43, 3424, 717}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("冒泡排序后的数组是: ", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test55(arr *[5]int, leftIndex int, rightIndex int, findNum int) {
	// 二分查找法
	// 将数组进行从小到大排序
	test56(arr)
	fmt.Println("从小到大排序后的数组是: ", arr)
	test57(arr, leftIndex, rightIndex, findNum)
	// 分割线
	println("-------------分割线--------------")
}

func test56(arr *[5]int) {
	//将数组进行从小到大排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func test57(arr *[5]int, leftIndex int, rightIndex int, findNum int) {
	// 二分查找法
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
		return
	}
	var midIndex = (leftIndex + rightIndex) / 2
	if (*arr)[midIndex] == findNum {
		fmt.Printf("找到了，下标是：%v \n", midIndex)
	} else if (*arr)[midIndex] > findNum {
		test57(arr, leftIndex, midIndex-1, findNum)
	} else {
		test57(arr, midIndex+1, rightIndex, findNum)
	}
}
