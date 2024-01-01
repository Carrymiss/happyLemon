package domain

import (
	"encoding/json"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *Monster) Store() bool {
	// 将monster序列化
	data, err := json.Marshal(m)
	if err != nil {
		println("序列化失败 err=", err)
		return false
	}

	// 保存到文件
	path := "C:/Users/Rei/Desktop/3.txt"
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		println("保存文件失败 err=", err)
		return false
	}
	return true
}

func (m *Monster) ReStore() bool {
	// 从文件读取
	path := "C:/Users/Rei/Desktop/3.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		println("读取文件失败 err=", err)
		return false
	}

	//反序列化
	err = json.Unmarshal(data, m)
	if err != nil {
		println("反序列化失败 err=", err)
		return false
	}
	return true
}
