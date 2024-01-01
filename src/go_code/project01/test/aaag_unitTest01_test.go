package test

import "testing"

// 被测试函数
func test104(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func Test1(t *testing.T) {
	res := test104(10)
	if res != 56 {
		t.Fatalf("test104(10) error, expected:%v, actual:%v", 56, res)
	}
	// 如果正确，输出日志
	t.Logf("test104(10) success, expected:%v, actual:%v", 56, res)
	// 分割线
	println("-------------分割线--------------")
}
