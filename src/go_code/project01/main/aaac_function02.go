package main

import (
	"fmt"
	"happyLemon/src/go_code/project01/demo01"
	"strconv"
	"strings"
	"time"
)

// 定义全局变量
var age = test19()

// 全局匿名函数
var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// init函数
	test18()

	// 匿名函数
	test20()
	test21()
	test22()

	// 闭包
	test23()
	test26()

	// defer(延时机制) 关键字
	test28()

	// 字符串函数
	test29()

	// 时间日期常用函数
	test30()
}

// init函数
func init() {
	demo01.Age = 100
	demo01.Name = "jack"
	fmt.Println("init函数执行了")
}

func test19() int {
	// 定义全局变量
	fmt.Println("全局变量执行了")
	return 90
}

func test18() {
	fmt.Println("main函数执行了")
	fmt.Println("age=", age)
	fmt.Println("demo01.Age=", demo01.Age)
	fmt.Println("demo01.Name=", demo01.Name)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test20() {
	// 匿名函数
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test21() {
	// 匿名函数
	// 将匿名函数赋值给一个变量（函数变量），则可以通过该变量来调用该匿名函数
	a := func(n1 int, n2 int) int { // 将匿名函数赋值给变量a 类型是myFunType 自定义类型 func(int, int) int
		return n1 + n2
	}
	res := a(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("--------------分割线--------------")
}

func test22() {
	res := fun1(10, 20)
	fmt.Println("res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test23() {
	// 闭包
	// 闭包是一个函数，这个函数包含了他外部作用域的一个变量
	// 底层原理：函数和他外部变量的引用（环境）
	// 闭包的说明：
	// 1.返回的是一个函数，但是这个函数引用到函数外的变量
	// 2.这个返回的函数在他被调用时，仍然可以找到他引用的外部变量
	// 3.闭包函数常见的错误用法：返回一个局部变量的地址
	// 4.闭包函数的价值：可以保存一个状态，可以反复使用这个状态（闭包函数是一个引用类型）
	f := test24()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))
	// 分割线
	fmt.Println("-------------分割线--------------")
}

// 累加器
func test24() func(int) int {
	var num int
	return func(x int) int {
		num += x
		return num
	}
}

// 累加器
func test25() func() {
	var str = "hello"
	return func() {
		str += string(36)
		fmt.Println("str=", str)
	}
}

func test26() {
	// 闭包
	f := test25()
	f()
	f()
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test27(num1 int, num2 int) int {
	// defer(延时机制) 关键字 用于函数结束之前释放资源 一般用于关闭文件或者数据库连接
	// defer的执行顺序是先进后出 defer的参数会在声明时就确定
	// defer将语句放入到栈时，也会将相关的值拷贝同时入栈
	defer fmt.Println("num1=", num1) // 最后执行
	defer fmt.Println("num2=", num2) // 然后执行
	res := num1 + num2
	fmt.Println("res=", res) // 先执行
	return res
}

func test28() {
	// defer(延时机制) 关键字 用于函数结束之前释放资源 一般用于关闭文件或者数据库连接
	res := test27(10, 20)
	fmt.Println("main_res=", res)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

func test29() {
	// len函数 用于求长度，比如string、array、slice、map、channel
	str := "hello北" // golang的编码统一为utf-8，一个汉字占3个字节
	fmt.Println("str_len1=", len(str))
	str = "hello"
	fmt.Println("str_len2=", len(str))
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 字符串遍历
	str = "hello北京"
	str1 := []rune(str) // 将str转换成[]rune切片
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c \n", str1[i])
	}
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 字符串转成整数
	num, err := strconv.Atoi("121345")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果num=", num)
	}
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v str的类型是：%T \n", str, str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 字符串转byte切片
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%c \n", bytes)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// byte切片转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v \n", str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 10进制转成其他进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("对应的2进制是：%v \n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("对应的16进制是：%v \n", str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 查找子串是否在指定的字符串中
	str2 := strings.Contains("seafood", "foo")
	fmt.Printf("是否存在对应的子串: %v \n", str2)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 统计一个字符串有几个指定的子串
	num = strings.Count("ceheese", "e")
	fmt.Printf("字符串中有%v个e \n", num)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 不区分大小写的字符串比较（==是区分字母大小写的）
	str2 = strings.EqualFold("abc", "Abc")
	fmt.Printf("不区分大小写的字符串比较结果是：%v \n", str2)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 返回子串在字符串第一次出现的index值，如果没有返回-1
	num = strings.Index("NLT_abc", "abc")
	fmt.Printf("子串在字符串第一次出现的index值是：%v \n", num)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 返回子串在字符串最后一次出现的index值，如果没有返回-1
	num = strings.LastIndex("go golang", "go")
	fmt.Printf("子串在字符串最后一次出现的index值是：%v \n", num)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 将指定的子串替换成另外一个子串，n表示替换几个，如果n=-1表示全部替换
	str = strings.Replace("go go hello", "go", "go语言", -1)
	fmt.Printf("替换后的字符串是：%v \n", str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("strArr[%v]=%v \n", i, strArr[i])
	}
	fmt.Printf("strArr=%v \n", strArr)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 将字符串的字母进行大小写的转换
	str = "goLang Hello"
	str4 := strings.ToLower(str)
	str3 := strings.ToUpper(str)
	fmt.Printf("str4=%v str3=%v \n", str4, str3)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 将字符串左右两边的空格去掉
	str = strings.TrimSpace(" tn a lone gopher ntrn ")
	fmt.Printf("str=%q \n", str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 将字符串左右两边指定的字符去掉
	str = strings.Trim("! hello! ", " !")
	fmt.Printf("str=%q \n", str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 将字符串左边指定的字符去掉
	str = strings.TrimLeft("! hello! ", " !")
	fmt.Printf("str=%q \n", str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 将字符串右边指定的字符去掉
	str = strings.TrimRight("! hello! ", " !")
	fmt.Printf("str=%q \n", str)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 判断字符串是否以指定的字符串开头
	str2 = strings.HasPrefix("ftp://192.168.75.1", "ftp")
	fmt.Printf("str2=%v \n", str2)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 判断字符串是否以指定的字符串结尾
	str2 = strings.HasSuffix("ftp://192.168.75.1", "75.1")
	fmt.Printf("str2=%v \n", str2)
	// 分割线
	fmt.Println("-------------分割线--------------")
}

type Duration int64

func test30() {
	// 时间日期常用函数
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now的类型是：%T \n", now, now)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 获取年月日时分秒
	fmt.Printf("年=%v \n", now.Year())
	fmt.Printf("月=%v \n", now.Month())
	fmt.Printf("日=%v \n", now.Day())
	fmt.Printf("时=%v \n", now.Hour())
	fmt.Printf("分=%v \n", now.Minute())
	fmt.Printf("秒=%v \n", now.Second())
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 格式化日期时间的第二种方式
	fmt.Printf("%s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("%s\n", now.Format("2006-01-02"))
	fmt.Printf("%s\n", now.Format("15:04:05"))
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 时间常量
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	// 分割线
	fmt.Println("-------------分割线--------------")

	// 时间戳
	fmt.Printf("unix时间戳=%v unix时间戳的类型是：%T \n", now.Unix(), now.Unix())
	fmt.Printf("unixnano时间戳=%v unixnano时间戳的类型是：%T \n", now.UnixNano(), now.UnixNano())
	// 分割线
	fmt.Println("-------------分割线--------------")
}
