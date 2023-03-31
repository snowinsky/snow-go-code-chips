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
func NewA() *A {
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

func DigDemo() {
	container := dig.New()
	container.Provide(NewB)
	container.Provide(NewA)
	container.Provide(NewC)
	container.Provide(NewD)

	container.Invoke(func(b *B, c *C, d *D) {
		log.Println(b, c, d)
	})

}
