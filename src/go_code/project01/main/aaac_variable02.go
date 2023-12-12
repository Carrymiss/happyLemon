package main

import "fmt"

func main() {
	// 切片
	test45()
	test46()
	test47()
	test48()
	test49()
	test50()
	test51()
	test53()
	test52()
}

func test45() {
	// 切片的英文是slice
	// 切片是数组的一个引用，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	// 切片的定义：var 切片名 []类型
	// 切片的初始化：var 切片名 []类型 = []类型{元素1,元素2,...}
	// 切片的长度是可以变化的
	var arr1 = [5]int{1, 22, 55, 66, 77}
	// 定义一个切片 左闭右开 也就是说包含下标为1的元素，不包含下标为3的元素
	var slice1 = arr1[1:3]
	fmt.Println("当前切片的值是：", slice1)
	fmt.Println("当前切片的元素个数是：", len(slice1))
	// 切片的容量是指底层数组的容量 是动态变态化的
	fmt.Println("当前切片的容量是：", cap(slice1))
	// 分割符
	println("-------------分割线--------------")

	fmt.Printf("arr[1]的地址是：%p \n", &arr1[1])
	fmt.Printf("arr[2]的地址是：%p \n", &arr1[2])
	// TODO 这里和老师讲的不一样， 老师这边是切片是个结构体 由三部分构成 地址(指向的第一个元素)、长度(根据指定计算出来)和容量(自动进行伸缩)
	fmt.Printf("当前切片的地址是：%p,第一个元素的地址是：%p,第二个元素的地址是：%p \n", &slice1, &slice1[0], &slice1[1])
	// 分割符
	println("-------------分割线--------------")

	slice1[0] = 44
	fmt.Println("当前切片的值是：", slice1)
	fmt.Println("当前数组的值是：", arr1)
	// 分割符
	println("-------------分割线--------------")
}

func test46() {
	// 第二种创建切片的方式
	// var 切片名 []类型 = make([]类型,长度,容量)
	var slice2 = make([]int, 5, 10)
	fmt.Println("当前切片的值是：", slice2)
	fmt.Println("当前切片的元素个数是：", len(slice2))
	fmt.Println("当前切片的容量是：", cap(slice2))
	// 分割符
	println("-------------分割线--------------")
}

func test47() {
	// 第三种创建切片的方式
	// 直接指定数组的值，类似与make的方式
	// 此时的容量和长度相等
	// var 切片名 []类型 = []类型{元素1,元素2,...}
	var slice3 = []int{1, 2, 3, 4, 5}
	fmt.Println("当前切片的值是：", slice3)
	fmt.Println("当前切片的元素个数是：", len(slice3))
	fmt.Println("当前切片的容量是：", cap(slice3))
	// 分割符
	println("-------------分割线--------------")
}

func test48() {
	// 切片的遍历
	var slice3 = []int{1, 2, 3, 4, 5}
	for index, value := range slice3 {
		fmt.Printf("当前切片的下标是：%v，值是：%v \n", index, value)
	}
	// 分割符
	println("-------------分割线--------------")
}

func test49() {
	// 切片可以继续切片
	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = slice3[1:]
	fmt.Println("当前切片slice4的值是：", slice4)
	var slice5 = slice4[:2]
	fmt.Println("当前切片slice5的值是：", slice5)
	// 分割符
	println("-------------分割线--------------")
}

func test50() {
	// 给切片追加元素
	var slice3 = []int{1, 2, 3, 4, 5}
	slice3 = append(slice3, 6, 7, 8)
	fmt.Println("当前切片slice3的值是：", slice3)
	// 分割符
	println("-------------分割线--------------")

	// 追加切片
	var slice4 = []int{11, 22, 33}
	// 这三个点必须带上
	slice3 = append(slice3, slice4...)
	fmt.Println("当前切片slice3的值是：", slice3)
	// 分割符
	println("-------------分割线--------------")
}

func test51() {
	// 切片的拷贝
	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, 10)
	copy(slice4, slice3) // 将slice3拷贝到slice4中
	fmt.Println("当前切片slice4的值是：", slice4)
	// 分割符
	println("-------------分割线--------------")
}

func test53() {
	// string的底层是一个byte数组，因此string也可以进行切片处理
	var str = "hello world"
	var slice1 = str[1:3]
	fmt.Println("当前切片slice1的值是：", slice1)
	// 分割符
	println("-------------分割线--------------")
}

func test52() {
	// 如果需要修改字符串，可以先将其转换成[]rune或[]byte，进行修改，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	// 将字符串转换成[]byte类型
	var str = "hello world"
	var slice1 = []byte(str) // 将字符串转换成[]byte类型
	slice1[0] = 'H'
	// 将[]byte转换成string类型
	str = string(slice1)
	fmt.Println("当前字符串str的值是：", str)
	// 分割符
	println("-------------分割线--------------")

	// 将字符串转换成[]rune类型 可以处理中文字符
	var str2 = "hello world"
	var slice2 = []rune(str2) // 将字符串转换成[]rune类型
	slice2[0] = '你'
	// 将[]rune转换成string类型
	str2 = string(slice2)
	fmt.Println("当前字符串str2的值是：", str2)
	// 分割符
	println("-------------分割线--------------")
}
