package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 文件操作
	test90()
	test91()
	test92()
}

func test90() {
	// 打开文件
	file, err := os.Open("C:/Users/Rei/Desktop/1.txt")
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}

	// 读取文件
	var buf [128]byte
	n, err := file.Read(buf[:])
	if err != nil {
		println("读取文件失败=", err)
	}
	fmt.Println("读取的字节数=", n)
	fmt.Println("读取的内容=", string(buf[:n]))
	// 分割符
	fmt.Println("-------------分割线--------------")

	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("关闭文件失败=", err)
		}
	}(file)
}

// TODO 使用  ReadString 无法读取到最后一行
func test91() {
	// 打开文件
	file, err := os.Open("C:/Users/Rei/Desktop/1.txt")
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}

	// 读取文件
	// 创建一个*Reader 默认缓冲区为4096
	reader1 := bufio.NewReader(file)
	// 循环读取文件
	for {
		readString, err := reader1.ReadString('\n') // 读到换行就结束
		// 判断是否读到文件末尾 （io.EOF等于读到文件的末尾）
		if err == io.EOF {
			break
		}
		fmt.Print("读取的内容=", readString)
	}
	fmt.Println("读取文件结束")
	// 分割符
	fmt.Println("-------------分割线--------------")

	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)
}

func test92() {
	// 文件的打开和关闭都封装到ReadFile()函数中
	// 使用io包下的ioutil.ReadFile()一次性将文件读取到位
	file := "C:/Users/Rei/Desktop/1.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("读取文件失败=", err)
	}
	fmt.Println("读取的内容=", string(content))
	// 分割符
	fmt.Println("-------------分割线--------------")
}

func test93() {
	// 打开文件
	file, err := os.OpenFile("C:/Users/Rei/Desktop/1.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}

	// 写入文件
	// 创建一个*Writer 默认缓冲区为4096
	writer1 := bufio.NewWriter(file)
	// 写入文件
	writer1.WriteString("hello world\r\n")
	writer1.WriteString("hello golang\r\n")
	// 将缓冲区的数据写入文件
	writer1.Flush()
	fmt.Println("写入文件结束")
	// 分割符
	fmt.Println("-------------分割线--------------")

	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)
}
