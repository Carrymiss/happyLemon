package domain

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

// 工厂函数 相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age和sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		// fmt.Println("年龄范围不正确")
		// 我们这里可以给一个默认值
		p.age = 0
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		// fmt.Println("薪水范围不正确")
		// 我们这里可以给一个默认值
		p.sal = 0.0
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}

type account struct {
	// 账号 长度在6-10位
	AccountNo string
	// 密码 必须是6位
	Password string
	// 余额 必须大于20
	Balance float64
}

// 工厂函数 相当于构造函数
func NewAccount(accountNo string, password string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不正确")
		return nil
	}
	if len(password) != 6 {
		fmt.Println("密码长度不正确")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额数目不正确")
		return nil
	}
	return &account{
		AccountNo: accountNo,
		Password:  password,
		Balance:   balance,
	}
}
