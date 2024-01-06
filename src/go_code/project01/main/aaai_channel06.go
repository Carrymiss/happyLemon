package main

func main() {
	// 管道可以声明为只读或者只写
	// 1.在默认情况下，管道是双向的
	// var chan1 chan int  可读可写
	// 2.声明为只写
	test116()

}
func test116() {
	// 声明为只写
	var chan1 chan<- int
	chan1 = make(chan int, 3)
	chan1 <- 20
	// num := <-chan1 // error
	// println("num=", num)

	// 分割线
	println("-------------分割线--------------")

	// 声明为只读
	var chan2 <-chan int
	num2 := <-chan2
	println("num2=", num2)
	// 分割线
	println("-------------分割线--------------")
}
