package main

import "fmt"

func main() {
	// 数组
	test37()
	test38()
	test39()
	test40()
	// 数组是值拷贝
	test41()
	test43()
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
	fmt.Printf("数组的地址是：%p,第一个元素的地址是：%p,第二个元素的地址是：%p \n", &arr, &arr[0], &arr[1])
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

func test40() {
	// 数组的遍历
	var arr2 = [3]int{1, 2, 3}
	for index, value := range arr2 {
		fmt.Printf("数组的下标是 %v,该下标的值是 %v \n", index, value)
	}
	var arr4 = [4]string{1: "tom", 2: "jack", 3: "marry"}
	for _, value := range arr4 {
		fmt.Printf("数组的值是 %v\n", value)
	}
	// 分割线
	println("-------------分割线--------------")
}

func test41() {
	var arr = [3]int{1, 2, 3}
	// 因为是不同的方法栈 而且数组是值拷贝 所以不会改变原数组的值
	test42(arr)
	fmt.Println("arr=", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test42(arr [3]int) {
	arr[0] = 10
}

func test43() {
	var arr = [3]int{1, 2, 3}
	// 拿到数组的首地址 并输出当前的值
	fmt.Printf("拿到数组的首地址是：%p,此时的值是 %v \n", &arr, arr)
	test44(&arr)
	fmt.Println("arr=", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test44(arr *[3]int) {
	(*arr)[0] = 10
	fmt.Printf("拿到数组的首地址是：%p,此时的值是 %v \n", &arr, arr)
}
