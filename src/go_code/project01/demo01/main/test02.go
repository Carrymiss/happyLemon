package main

import (
	"fmt"
	"strings"
)

func main() {
	// 如果name没有后缀，就加上后缀，否则就返回原来的名字
	test22()
}

func test21(suffix string) func(string) string {
	return func(name string) string {
		// 如果name没有后缀，就加上后缀，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func test22() {
	f := test21(".jpg")
	name1 := f("test")
	name2 := f("bird.jpg")
	fmt.Println("文件名处理后：", name1)
	fmt.Println("文件名处理后：", name2)
	// 分割线
	println("-------------分割线--------------")
}
