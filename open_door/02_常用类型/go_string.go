package main

import (
	"fmt"
	"strconv"
)

func main() {
	//变量声明和赋值 两个变量可以一起声明赋值
	var p1, p2 string = "abc", "efg"
	fmt.Println(p1, p2)

	//不给初始值的时候就会有默认值 string的默认值不是nil而是空
	//不声明变量类型但是赋值了的话，编译器会自行推断变量类型
	var p3 string
	fmt.Println("p3=" + p3)

	//下划线，空变量占位符，如果某个变量不得不取出来，但又不用的话，就可以用占位符取出来。这样后面的程序是不能引用它的值的
	p7, _, p9 := true, "abc", 123
	fmt.Println(p7, p9)

	var digitStr1 = "1545"
	var digitStr2 = "0.1"
	var digitStr3 = "-3.1415926"

	digit1, _ := strconv.Atoi(digitStr1)
	fmt.Println(digit1)
	digit2, _ := strconv.ParseFloat(digitStr2, 8)
	fmt.Println(digit2)
	digit3, _ := strconv.ParseFloat(digitStr3, 16)
	fmt.Println(digit3)

	digitInt := 2343242
	var digitFloat float64 = 3.1415665465
	var digitIntStr = strconv.Itoa(digitInt)
	fmt.Println(digitIntStr)
	//将浮点数转成字符串，且本色还原的效果就是这样的。只支持float64位的数，32位的数必须先转成64的再转，否则会有精度变化
	var digitFloatStr = strconv.FormatFloat(digitFloat, 'f', 10, 64)
	fmt.Println(digitFloatStr)
}
