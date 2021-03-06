package demo

import "fmt"

// 定义一个抽象的组件
type Component interface {
	Operate()
}

// 实现一个具体的组件：组件1
type Component1 struct{}

func (c *Component1) Operate() {
	fmt.Println("component1 operate")
}

// 定义一个抽象的装饰者
type Decorator interface {
	Component // 继承抽象组件(其实就是接口组合的过程)
	Do()      // 这个是有额外的方法
}

// 实现一个具体的装饰者
type Decorator1 struct {
	Component Component
}

func (d *Decorator1) Operate() {
	d.Do()
	d.Component.Operate()
}

func (d *Decorator1) Do() {
	fmt.Println("发生了装饰行为")
}
