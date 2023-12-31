package model

import "fmt"

// 声明一个结构体 表示客户信息
type Customer struct {
	Id     int
	Name   string
	Gendre string
	Age    int
	Phone  string
	Email  string
}

// NewCustomer 使用工厂模式,返回一个Customer的实例
func NewCustomer(id int, name string, gendre string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gendre: gendre,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// GetInfo 返回客户信息
func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", c.Id, c.Name, c.Gendre, c.Age, c.Phone, c.Email)
	return info
}
