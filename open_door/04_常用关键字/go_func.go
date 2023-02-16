package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(add(1, 3))
	printlnAdd(2, 4)
	fmt.Println(addWithMessage(1, 9))
	fmt.Println(addWithMessage(-1, -9))

	//用下划线来承接一个不想使用的变量，这个变量因为没有变量名可以引用，所以就没法使用
	var a, _ = addWithMessage(2, 1)
	fmt.Println(a)

	//函数也可以成为一种自定义类型
	//当指定函数的具体类型时，后面实现的函数体必须和指名的函数类型的结构相一致，入参，出参都必须一致
	var myFunc1Impl MyFunc1 = func(a string, b int) (bool, int) {
		return len(a) > b, len(a) - b
	}
	fmt.Println(myFunc1Impl("sss", 3))
	fmt.Println(myFunc1Impl.toFuncName())
	//如果不指名函数的具体类型，后面的函数体就可以随便写。而这种函数其实就是普通的代码片段，而且这个代码片段还是可以随便用的，而这就是所谓的闭包
	var myFunc2Impl = func(a string) bool {
		return true
	}
	fmt.Println(myFunc2Impl("a"))

}

// 最普通的函数，带入参和出参
func add(a int, b int) int {
	return a + b
}

// 无返回值的参数
func printlnAdd(a int, b int) {
	fmt.Println(a + b)
}

// 多个返回值的函数
func addWithMessage(a int, b int) (int, string) {
	if a > 0 && b > 0 {
		return a + b, "正整数加法"
	} else {
		return a + b, "非正整数加法"
	}
}

// 函数也是可以当参数传递的
type MyFunc1 func(a string, b int) (bool, int)

// func后面跟着括号的，是给类型增加新的方法，该方法则使用该类型的变量可以调用
func (p MyFunc1) toFuncName() string {
	return reflect.TypeOf(p).String()
}
