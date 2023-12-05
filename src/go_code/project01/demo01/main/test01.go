package main

import "fmt"

func main() {
	test01()
}

func test01() {
	// 假如还有97天放假，问：xx个星期零xx天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Println(week, "个星期零", day, "天")
}
