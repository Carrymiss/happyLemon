package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁
	lock sync.Mutex
)

// 计算n的阶乘 把结果放到map中
func test106(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 加锁
	lock.Lock()

	// TODO 这里没有互斥锁 会有并发问题 fatal error: concurrent map writes
	myMap[n] = res // 将结果放入到myMap中

	// 解锁
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go test106(i)
	}

	// 休眠10秒钟
	// 为了等上面的协程执行完毕
	time.Sleep(time.Second * 10)

	// 遍历map的值
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
