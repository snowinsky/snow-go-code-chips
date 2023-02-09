package main

import (
	"fmt"
)

// bool不能随意转化成其他类型，go的类型是非常严格的
func main() {
	//最常用，最标准的定义变量的方法。var+变量名+变量类型+=+赋初始值
	var isSuccess bool = true
	fmt.Println(isSuccess)
	//其中一种简化方式，var+变量名+=+赋初值，省略了变量类型，这是因为可以从赋的初始值上推断出变量类型
	var isFail = false
	fmt.Println(isFail)
	//这是先定义变量，var+变量名+变量类型，一个都不能省略；然后再在必要的地方给变量赋值
	var isEmpty bool
	isEmpty = true
	fmt.Println(isEmpty)
	//另外一种变量定义方式，可以同时省略掉var和变量类型两个元素，使用:=来代替。使用:=之后表示左边的是变量名，右边的是初始值。最简单的方式了
	hasNext := false
	fmt.Println(hasNext)

	//常量的定义方式
	const isSuccessDefaultValue bool = false
	fmt.Println(isSuccessDefaultValue)

}
