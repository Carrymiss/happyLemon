package main

import "happyLemon/src/go_code/project01/demo3/model"

func main() {
	// 运行项目
	customreView := model.CustomerView{
		Key:  "",
		Loop: true,
	}
	customreView.MainMenu()
}
