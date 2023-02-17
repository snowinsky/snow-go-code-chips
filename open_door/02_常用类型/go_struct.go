package main

import "fmt"

type Color struct {
	id   int
	name string
}

func (c Color) draw() {
	fmt.Println("draw the color is " + c.name)
}

func main() {
	//struct就是javabean
	red := new(Color)
	fmt.Println(red)
	fmt.Println(red.name)
	fmt.Println(red.id)
	red.draw()

	blue := Color{1, "blue"}
	fmt.Println(blue)
	blue.draw()

}
