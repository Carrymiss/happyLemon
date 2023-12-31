package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 求2数和
	test62()

	// 判断输入的参数是什么类型
	n1 := 1.1
	n2 := 2.2
	n3 := 30
	n4 := "hello"
	n5 := true
	n6 := Calcuator{1.2, 2.3}
	test63(n1, n2, n3, n4, n5, n6, &n6)

	// 创建一个新文件 打印五次 hello Gardon
	test64()

	// 打开一个存在的文件 重新弄定义原来的内容
	test65()

	// 追加内容
	test66()

	// 先读后追加
	test67()

	// 将一个文件的内容写入到另一个文件中
	test68()
}

type Calcuator struct {
	num1 float64
	num2 float64
}

func (cal *Calcuator) test61() float64 {
	// 求和
	return cal.num1 + cal.num2
}

func test62() {
	var cal Calcuator
	cal.num1 = 1.2
	cal.num2 = 2.3
	cal.test61()
	fmt.Printf("两数的和是：%v\n", fmt.Sprintf("%.2f", cal.test61()))
	// 分割线
	println("-------------分割线--------------")
}

func test63(items ...interface{}) {
	// 判断输入的参数是什么类型
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index+1, v)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index+1, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index+1, v)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index+1, v)
		case Calcuator:
			fmt.Printf("第%v个参数是Calcuator类型，值是%v\n", index+1, v)
		case *Calcuator:
			fmt.Printf("第%v个参数是*Calcuator类型，值是%v\n", index+1, v)
		default:
			fmt.Printf("第%v个参数是未知类型，值是%v\n", index+1, v)
		}
	}
	// 分割线
	println("-------------分割线--------------")
}

func test64() {
	// 创建一个新文件 打印五次 hello Gardon
	// permission:0666 代表文件的权限 0666代表可读可写 0777代表可读可写可执行 0代表可读可写可执行 4代表可读 2代表可写 1代表可执行 6代表可读可写
	file, err := os.OpenFile("C:/Users/Rei/Desktop/2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)

	// 准备写入5句 hello Gardon
	str := "hello Gardon\n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(str)
	}
	// 因为writer是带缓存的，因此在调用WriterString方法时，其实内容是先写入到缓存的，所以需要调用Flush方法，将缓存的数据真正写入到文件中，否则文件中会没有数据
	_ = writer.Flush()
	println("写入完成")
	// 分割线
	println("-------------分割线--------------")
}

func test65() {
	// 打开一个存在的文件 重新弄定义原来的内容
	file, err := os.OpenFile("C:/Users/Rei/Desktop/2.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)

	// 准备写入5句 hello Gardon
	// 换行有些编译器是\r 有些是\n
	str := "hello World!\r\n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, _ = writer.WriteString(str)
	}
	// 因为writer是带缓存的，因此在调用WriterString方法时，其实内容是先写入到缓存的，所以需要调用Flush方法，将缓存的数据真正写入到文件中，否则文件中会没有数据
	_ = writer.Flush()
	println("写入完成")
	// 分割线
	println("-------------分割线--------------")
}

func test66() {
	// 追加内容
	file, err := os.OpenFile("C:/Users/Rei/Desktop/2.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)

	// 准备写入5句 hello Gardon
	// 换行有些编译器是\r 有些是\n
	str := "追加内容！！!\r\n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, _ = writer.WriteString(str)
	}
	// 因为writer是带缓存的，因此在调用WriterString方法时，其实内容是先写入到缓存的，所以需要调用Flush方法，将缓存的数据真正写入到文件中，否则文件中会没有数据
	_ = writer.Flush()
	println("写入完成")
	// 分割线
	println("-------------分割线--------------")
}

func test67() {
	// 先读后追加内容
	path := "C:/Users/Rei/Desktop/2.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)

	// 读取文件内容
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件失败=", err)
	}
	fmt.Println("读取的内容=", string(content))

	// 准备写入5句 hello Gardon
	// 换行有些编译器是\r 有些是\n
	str := "第二次追加内容！！!\r\n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, _ = writer.WriteString(str)
	}
	// 因为writer是带缓存的，因此在调用WriterString方法时，其实内容是先写入到缓存的，所以需要调用Flush方法，将缓存的数据真正写入到文件中，否则文件中会没有数据
	_ = writer.Flush()
	println("写入完成")
	// 分割线
	println("-------------分割线--------------")
}

// TODO 这个写入直接将原来的内容覆盖了
func test68() {
	// 将一个文件的内容写入到另一个文件中 将1.txt的内容写入到2.txt中
	path1 := "C:/Users/Rei/Desktop/1.txt"
	path2 := "C:/Users/Rei/Desktop/2.txt"

	// 读取文件内容
	data, err1 := ioutil.ReadFile(path1)
	if err1 != nil {
		fmt.Println("读取文件失败=", err1)
		return
	}

	err2 := ioutil.WriteFile(path2, data, 0666)
	if err2 != nil {
		fmt.Println("写入文件失败=", err2)
		return
	}
}
