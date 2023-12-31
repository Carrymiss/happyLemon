package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件操作
	// test90()
	test91()
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
