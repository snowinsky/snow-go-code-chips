package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type DrawService interface {
	draw() string
	drawFrom(a int) string
}

type StudentList []string

// panic相当于throw new RuntimeException
func (s StudentList) drawFrom(a int) string {
	//panic("implement me")
	return reflect.TypeOf(s).String() + " drawing...." + strconv.Itoa(a)
}

func (s StudentList) draw() string {
	return reflect.TypeOf(s).String() + " drawing....nothing"
}

func (s StudentList) length() int {
	return len(s)
}

func main() {
	//用接口类型定义变量时，如果实现类没有实现接口的所有方法，会报错的，并要求必须实现所有方法
	var ds DrawService = StudentList(make([]string, 5))
	fmt.Println(ds)
	fmt.Println(ds.draw())
	ds.drawFrom(1)

	//因为ds是接口类型，所以只能访问接口的方法，没法调用StudentList本身的方法
	//ds.length()
	//将接口强转为某一种具体类型, 就可以调用这种类型下的独有的方法了
	var list StudentList = ds.(StudentList)
	fmt.Println(list.length())

}
