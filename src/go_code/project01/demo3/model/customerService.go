package model

// CustomerService 完成对customer的增删改查
type CustomerService struct {
	customers []Customer
	// 声明一个字段,表示当前切片含有多少个客户
	customerNum int
}

// NewCustomerService 返回一个 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中,我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := NewCustomer(1, "张三", "男", 20, "110", "11231188@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// List 返回客户切片
func (cs *CustomerService) List() []Customer {
	return cs.customers
}

// Add 添加客户到customers切片
func (cs *CustomerService) Add(customer Customer) bool {
	// 我们确定一个分配id的规则,就是添加的顺序
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// Delete 根据id删除客户(从切片中删除)
func (cs *CustomerService) Delete(id int) bool {
	// 先找到该客户对应的下标
	index := cs.FindById(id)
	if index == -1 {
		return false
	}
	// 删除切片中的一个元素
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

// FindById 先找到该客户对应的下标
func (cs *CustomerService) FindById(id int) int {
	index := -1
	// 遍历切片
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
		}
	}
	return index
}
