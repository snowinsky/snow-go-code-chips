package main

import "fmt"

type Table struct {
	id   int
	name string
}

// 操作指针和操作指针的实际变量是一样的
func main() {
	var t1 Table = Table{1, "T_1"}
	fmt.Println(t1)
	//变量前面加&就会返回当前这个变量的指针，类型就是实际类型前面加*
	fmt.Println(&t1)
	var t2Point *Table = &Table{1, "T_2"}
	fmt.Println(t2Point)
	//变量前面加*就会返回当前这个指针变量的实际值，实际值得类型就是不带*的类型
	fmt.Println(*t2Point)
}
