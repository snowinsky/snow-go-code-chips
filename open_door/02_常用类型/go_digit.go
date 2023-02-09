package main

import "fmt"

// 默认值 只定义变量，不赋值的话，会被赋默认值
func main() {
	//数字类型 go的数字类型好复杂，挺多的
	var uintA uint = 1
	var uint8A uint8 = 255
	var uint16A uint16 = 65535
	var uint32A uint32 = 4294967295
	var uint64A uint64 = 18446744073709551615

	var int8A int8 = -128
	var int16A int16 = -32768
	var int32A int32 = -2147483648
	var int64A int64 = -9223372036854775808
	var runeA rune = -2147483648

	fmt.Println(uintA)
	fmt.Println(uint8A)
	fmt.Println(uint16A)
	fmt.Println(uint32A)
	fmt.Println(uint64A)
	fmt.Println(int8A)
	fmt.Println(int16A)
	fmt.Println(int32A)
	fmt.Println(int64A)
	fmt.Println(runeA)
}
