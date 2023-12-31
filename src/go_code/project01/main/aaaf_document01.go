package main

import "os"

func main() {
	test90()
}

func test90() {
	// 打开文件
	file, err := os.Open("C:/Users/Rei/Desktop/1.txt")
	if err != nil {
		println("打开文件失败=", err)
	}

	// 读取文件
	var buf [128]byte
	n, err := file.Read(buf[:])
	if err != nil {
		println("读取文件失败=", err)
	}
	println("读取的字节数=", n)
	println("读取的内容=", string(buf[:n]))

	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("关闭文件失败=", err)
		}
	}(file)
}
