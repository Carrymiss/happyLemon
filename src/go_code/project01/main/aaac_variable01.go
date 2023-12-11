package main

import "fmt"

func main() {
	// 数组
	test37()
	test38()
	test39()
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

func test38() {
	// 数组的遍历
	var arr [5]int
	fmt.Println("数组的默认值是：", arr)
	// 这里int64占用8个字节，所以个元素的地址相差8个字节，如果是int32则相差4个字节
	fmt.Printf("数组的地址是：%p,第一个元素的地址是：%p,第二个元素的地址是：%p /n", &arr, &arr[0], &arr[1])
	// 分割线
	println("-------------分割线--------------")
}

func test39() {
	// 几种初始化数组的方式
	var arr1 [5]int
	var arr2 = [3]int{1, 2, 3}
	var arr3 = [...]int{1, 2, 3}
	// 可以指定元素对应的下标
	var arr4 = [4]string{1: "tom", 2: "jack", 3: "marry"}

	fmt.Println("arr1=", arr1)
	fmt.Println("arr2=", arr2)
	fmt.Println("arr3=", arr3)
	fmt.Println("arr4=", arr4)
	// 分割线
	println("-------------分割线--------------")
}
