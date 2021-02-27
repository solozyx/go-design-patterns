package demo

import (
	"fmt"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	p := &product1{}
	p.SetName("产品1")
	fmt.Println(p.GetName())

	p2 := &product2{}
	p2.SetName("产品2")
	fmt.Println(p2.GetName())
}

func TestProductFactory_Create(t *testing.T) {
	p := &ProductFactory{}
	product1 := p.Create(p1)
	product1.SetName("产品1")
	fmt.Println(product1.GetName())

	product2 := p.Create(p2)
	product2.SetName("产品2")
	fmt.Println(product2.GetName())
}
