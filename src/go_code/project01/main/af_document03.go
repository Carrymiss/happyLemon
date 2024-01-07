package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 统计不同类型的字符数量
	test95()
}

type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字个数
	SpaceCount int // 记录空格个数
	OtherCount int // 记录其他字符个数
}

func test95() {
	// 统计不同类型的字符数量
	path := "C:/Users/Rei/Desktop/1.txt"
	file, err1 := os.Open(path)
	if err1 != nil {
		fmt.Println("读取文件失败=", err1)
	}

	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败=", err)
		}
	}(file)

	reader := bufio.NewReader(file)
	var count CharCount

	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF 表示文件的末尾
			break
		}
		// 兼容中文字符
		str1 := []rune(str)
		// 遍历str，进行统计
		for _, v := range str1 {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其他字符的个数为=%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
