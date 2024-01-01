package test

import (
	"happyLemon/src/go_code/project01/demo01/domain"
	"testing"
)

func Test2(t *testing.T) {
	monster := &domain.Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "牛魔拳",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() error , expected:%v, actual:%v", true, res)
	}
	t.Logf("monster.Store() success, expected:%v, actual:%v", true, res)

	// 分割线
	println("-------------分割线--------------")
}

func Test3(t *testing.T) {
	var monster domain.Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() error , expected:%v, actual:%v", true, res)
	}
	t.Logf("monster.ReStore() success, expected:%v, actual:%v", true, res)
	// 分割线
	println("-------------分割线--------------")
}
