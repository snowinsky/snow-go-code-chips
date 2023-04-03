package api

import (
	"go.uber.org/dig"
	"log"
)

type A struct{}
type B struct{}

func NewB() *B {
	return new(B)
}
func NewA(a B) *A {
	log.Println(a)
	return new(A)
}
func NewC() *C {
	return new(C)
}
func NewD() *D {
	return new(D)
}

type C struct{}

type D struct{}

// 依赖注入dig
func DigDemo() {
	//创建一个容器
	container := dig.New()
	//将各个对象的构造函数装入容器, 他们之间的相互调用关系，依赖关系，都交给容器去梳理
	container.Provide(NewA)
	container.Provide(NewB)
	container.Provide(NewC)
	container.Provide(NewD)

	//使用的时候，直接使用指针引用就行了
	container.Invoke(func(a *A, b *B, c *C, d *D) {
		log.Println(a, b, c, d)
	})

}
