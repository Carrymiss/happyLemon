package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	test104()
	for i := 0; i < 10; i++ {
		fmt.Println("main hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test104() {
	for i := 0; i < 10; i++ {
		fmt.Println("test103() hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
