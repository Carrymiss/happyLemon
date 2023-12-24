package demo01

import "fmt"

var Age int
var Name string
var HeroName string = "超人"

func Test01(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误")
	}
	return res
}

type Person struct {
	Name string
	Age  int
}

type person1 struct {
	Name string
	Age  int
}

// 大写的是可以外部引用 小写的想要引用就需要用到工厂模式
func NewPerson(name string, age int) *person1 {
	return &person1{
		Name: name,
		Age:  age,
	}
}

func (per *person1) GetAge() int {
	// 这里因为编译器底层做了优化 本来是(*per).Age
	return per.Age
}
