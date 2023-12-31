package model

// 完成对customer的增删改查
type CustomerService struct {
	customers []Customer
	// 声明一个字段,表示当前切片含有多少个客户
	customerNum int
}
