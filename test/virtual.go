package main

import "fmt"

type IBase interface {
	EncodeBody() string
}

type IA interface {
	IBase
	ToString() string
}

type A struct {
	IBase
}

func NewA() *A {
	a := A{}
	a.IBase = &a
	return &a
}

func (p *A) ToString() string {
	return "A hello" + p.IBase.EncodeBody()
}
func (p *A) EncodeBody() string {
	return "A body"
}

type B struct {
	A
}

func NewB() *B {
	b := B{}
	b.IBase = &b
	return &b
}

func (p *B) EncodeBody() string {
	return "B body"
}

func main() {
	a := NewA()
	b := NewB()
	var i IA = a
	fmt.Println(i.ToString())
	i = b
	fmt.Println(i.ToString())
}
