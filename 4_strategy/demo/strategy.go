package demo

import "fmt"

// 实现一个上下文的类
type Context struct {
	Strategy
}

// 定义抽象的策略
type Strategy interface {
	Do()
}

// 实现具体的策略
type Strategy1 struct{}

func (s *Strategy1) Do() {
	fmt.Println("执行策略1")
}

type Strategy2 struct{}

func (s *Strategy2) Do() {
	fmt.Println("执行策略2")
}
