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

	// 有个升序好的数组，现在插入一个数，要求此数组依然是升序
	test49()

	// 定义一个3行4列的二维数组，编写程序将四周的数据清0
	test50()

	// 定义一个4行4列的二维数组，将二维数组的第一行和第四行进行交换，第二行和第三行进行交换
	test51()

	// 数组arr [10]string里面保存了10个元素，查找“AA”是否存在，输出其下标 可能存在多个“AA” 也要输出其下标
	var arr = [10]string{"AA", "BB", "CC", "AA", "AA", "AA", "AA", "AA", "AA", "AA"}
	test52(arr)

	// 查找当前数组中最大数和最小数，已经对应的下标
	test53()

	// 找出当前数组大于平均值以及小于平均值的个数
	test54()

	// 存放学生的姓名和性别 使用map
	test55()

	// 使用map[string]map[string]string的map类型
	// key 表示用户名，是唯一的，不可以重复
	// 如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息
	test56()
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
	// 随机生成10个整数(1-100的范围)保存到数组，并倒序打印以及求平均值，求最大值、最大值的下标，并查找里面是否有55
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
	test46(arr)
	// 求平均值
	test47(arr)
	// 查找里面是否有55
	test48(arr, 55)
}
func test45(arr [10]int) {
	// 倒序打印
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("倒叙打印数组为：", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test46(arr [10]int) {
	// 求最大值及其坐标
	var max int = arr[0]
	var maxIndex int = 0
	for index, value := range arr {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	fmt.Printf("最大值是：%v,最大值的下标是：%v \n", max, maxIndex)
	// 分割线
	println("-------------分割线--------------")
}

func test47(arr [10]int) float64 {
	// 求平均值
	var sum int = 0
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("平均值是：%v \n", float64(sum)/float64(len(arr)))
	// 分割线
	println("-------------分割线--------------")
	return float64(sum) / float64(len(arr))
}

func test48(arr [10]int, num int) {
	// 查找数组是否含有某个数 使用二分法
	// 将数组进行从小到大排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("从小到大排序后的数组是: ", arr)
	// 使用二分法查找
	var leftIndex int = 0
	var rightIndex int = len(arr) - 1
	var midIndex int = (leftIndex + rightIndex) / 2
	for leftIndex <= rightIndex {
		if arr[midIndex] == num {
			fmt.Printf("找到了，下标是：%v \n", midIndex)
			break
		} else if arr[midIndex] > num {
			rightIndex = midIndex - 1
			midIndex = (leftIndex + rightIndex) / 2
		} else {
			leftIndex = midIndex + 1
			midIndex = (leftIndex + rightIndex) / 2
		}
	}
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
	}
	// 分割线
	println("-------------分割线--------------")
}

func test49() {
	// 有个升序好的数组，现在插入一个数，要求此数组依然是升序
	var arr = [5]int{1, 2, 3, 4, 5}
	var num int = 3
	var index int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > num {
			index = i
			break
		}
	}
	var arr2 [6]int
	for i := 0; i < len(arr2); i++ {
		if i < index {
			arr2[i] = arr[i]
		} else if i == index {
			arr2[i] = num
		} else {
			arr2[i] = arr[i-1]
		}
	}
	fmt.Println("插入后的数组是：", arr2)
	// 分割线
	println("-------------分割线--------------")
}

func test50() {
	// 定义一个3行4列的二维数组，编写程序将四周的数据清0
	var arr [3][4]int
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	for index1, value1 := range arr {
		for index2, _ := range value1 {
			arr[index1][index2] = rand.Intn(100) + 1
		}
	}
	fmt.Println("此时生成的数组是：", arr)
	// 将四周的数据清0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
		}
	}
	fmt.Println("此时生成的数组是：", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test51() {
	// 定义一个4行4列的二维数组，将二维数组的第一行和第四行进行交换，第二行和第三行进行交换
	var arr [4][4]int
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	for index1, value1 := range arr {
		for index2, _ := range value1 {
			arr[index1][index2] = rand.Intn(100) + 1
		}
	}
	fmt.Println("此时生成的数组是：", arr)
	// 交换
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("此时生成的数组是：", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test52(arr [10]string) {
	// 数组arr [10]string里面保存了10个元素，查找“AA”是否存在，输出其下标 可能存在多个“AA” 也要输出其下标
	// 定义一个切片用来保存下标
	var indexes []int
	for index, value := range arr {
		if value == "AA" {
			indexes = append(indexes, index)
		}
	}
	fmt.Println("AA的下标是：", indexes)
	// 分割线
	println("-------------分割线--------------")
}

func test53() {
	// 查找当前数组中最大数和最小数，已经对应的下标
	var arr [10]int
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	for index, _ := range arr {
		arr[index] = rand.Intn(10000) + 1
	}
	fmt.Println("此时生成是数组是：", arr)
	// 求最大值及其坐标
	var max int = arr[0]
	var maxIndex int = 0
	for index, value := range arr {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	fmt.Printf("最大值是：%v,最大值的下标是：%v \n", max, maxIndex)
	// 求最小值及其坐标
	var min int = arr[0]
	var minIndex int = 0
	for index, value := range arr {
		if value < min {
			min = value
			minIndex = index
		}
	}
	fmt.Printf("最小值是：%v,最小值的下标是：%v \n", min, minIndex)
	// 分割线
	println("-------------分割线--------------")
}

func test54() {
	// 找出当前数组大于平均值以及小于平均值的个数
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for index, _ := range arr {
		arr[index] = rand.Intn(10000) + 1
	}
	fmt.Println("此时生成的数组是：", arr)
	// 求平均值
	var avgF float64 = test47(arr)
	// 大于平均数的数
	var moreTh int
	// 小于平均数的数
	var lessTh int
	for _, value := range arr {
		if float64(value) > avgF {
			moreTh++
		} else if float64(value) < avgF {
			lessTh++
		}
	}
	fmt.Printf("当前数组大于平均值的个数是：%v,当前数组小于平均值的个数是：%v \n", moreTh, lessTh)
	// 分割线
	println("-------------分割线--------------")
}

func test55() {
	// 存放学生的姓名和性别 使用map
	var stuMap = make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 2)
	stuMap["stu01"]["name"] = "tom"
	stuMap["stu01"]["sex"] = "男"
	stuMap["stu02"] = make(map[string]string, 2)
	stuMap["stu02"]["name"] = "jack"
	stuMap["stu02"]["sex"] = "女"
	fmt.Println("stuMap=", stuMap)
	// 分割线
	println("-------------分割线--------------")
}

func test56() {
	var userMap = make(map[string]map[string]string)
	userMap["user01"] = make(map[string]string, 2)
	userMap["user01"]["password"] = "123456"
	userMap["user01"]["name"] = "tom"
	userMap["user02"] = make(map[string]string, 2)
	userMap["user02"]["password"] = "123456"
	userMap["user02"]["name"] = "jack"
	fmt.Println("userMap=", userMap)
	test57(userMap, "tom", "123456")
	test57(userMap, "user02", "123456")
}

func test57(users map[string]map[string]string, username string, password string) {
	// 使用map[string]map[string]string的map类型
	// key 表示用户名，是唯一的，不可以重复
	// 如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息
	// 判断用户名是否存在
	_, fin := users[username]
	if fin {
		// 存在
		users[username]["password"] = "888888"
	} else {
		// 不存在
		users[username] = make(map[string]string, 2)
		users[username]["password"] = password
		users[username]["name"] = username
	}
	fmt.Println("userMap=", users)
	// 分割线
	println("-------------分割线--------------")
}
