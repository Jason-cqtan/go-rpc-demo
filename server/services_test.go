package main

import (
	"fmt"
	"testing"
)

func TestAddService(t *testing.T) {

}

func TestStruct(t *testing.T) {
	base := Base{}
	c := Concrete{
		Base: base,
	}
	c.SayHello()
}

type Base struct {
}

type Concrete struct {
	Base
}

type Concrete1 struct {
	*Base
}


func (b Base) SayHello() {
	fmt.Printf("Base say hello, " + b.Name())
}

func (b Base) Name() string {
	return "Base"
}

func (c Concrete) Name() string {
	return "Concrete"
}