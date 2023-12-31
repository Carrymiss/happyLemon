package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	test93()
}

func test93() {
	// 将一个文件拷贝到另一处文件夹中
	path1 := "C:/Users/Rei/Desktop/1.txt"
	path2 := "C:/Users/Rei/2.txt"
	_, err := test94(path1, path2)
	if err != nil {
		fmt.Println("拷贝文件失败=", err)
	} else {
		fmt.Println("拷贝文件成功")
	}
	// 分割线
	println("-------------分割线--------------")
}

func test94(path1 string, path2 string) (written int64, err error) {
	// 读取文件内容
	srcFile, err1 := os.Open(path1)
	if err1 != nil {
		fmt.Println("读取文件失败=", err1)
	}
	// 通过srcFile获取到	reader
	reader := bufio.NewReader(srcFile)

	file, err := os.OpenFile(path2, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败=", err)
	}

	writer := bufio.NewWriter(file)
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)
	return io.Copy(writer, reader)
}
