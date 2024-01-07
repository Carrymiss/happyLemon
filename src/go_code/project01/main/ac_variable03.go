package main

import "fmt"

func main() {
	// 二维数组
	test58()
	test59()
	test60()
}

func test58() {
	// 二维数组
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	arr[3][4] = 4
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			print(arr[i][j], " ")
		}
		println()
	}
	// 分割线
	println("-------------分割线--------------")

	// 正好差了48字节 也就是6个int64 因为上面定义的这个数组二维的元素有6个
	fmt.Printf("arr[0]的地址是：%p \n", &arr[0])
	fmt.Printf("arr[1]的地址是：%p \n", &arr[1])
	// 差了8个字节 也就是1个int64
	fmt.Printf("arr[1][0]的地址是：%p \n", &arr[1][0])
	fmt.Printf("arr[1][1]的地址是：%p \n", &arr[1][1])

	// 分割线
	println("-------------分割线--------------")
}

func test59() {
	// 创建二维数组的第二种方法
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr=", arr)

	arr1 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr1=", arr1)
	// 分割线
	println("-------------分割线--------------")
}

func test60() {
	// 二维数组遍历 for range方式
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for index, value := range arr {
		for index1, value1 := range value {
			fmt.Printf("arr[%v][%v]=%v ", index, index1, value1)
		}
		fmt.Println()
	}
	// 分割线
	println("-------------分割线--------------")
}
