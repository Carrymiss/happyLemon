package main

func main() {
	// 二维数组
	test58()
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
}
