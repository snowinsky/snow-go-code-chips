package main

import (
	"fmt"
	"strconv"
)

// 自定义类型，就是以已有类型为蓝本定义一个新的类型，这个新的类型只继承已有类型的底层结构，但不继承其方法
type myInt int

func (mi myInt) toString() string {
	return strconv.Itoa(int(mi))
}
func (mi myInt) toInt() int {
	return int(mi)
}

// 类型别名，就是给已有类型换了个名字
type anotherInt = string

//不能为非本地类型定义新方法，也就是说，这个类型的别名必须在定义它的包里边被定义新方法
//string所在的包是builtin，而现在的包是main，不能在main里定义string这个类型的别名下的新方法
//解决这个问题有两个方式
//1. 直接去builtin包中去为anotherInt来定义新方法
//2. 使用type anotherInt string的方式定义新类型，而不是使用类型别名
/*func (mi anotherInt) stringLength() int{
	return len(mi)
}*/

// 本包内定义的类型，就可以通过给别名添加方法的方式给别名所对应的本名添加方法
type Person struct {
	id   int
	name string
}

// 这是person的方法
func (p Person) sayName() string {
	return p.name
}

// 给person起个别名，添加个方法，其实就相当于给person添加了方法
type ChinesePerson = Person

func (cp ChinesePerson) sayChineseName() string {
	return "中文名——" + cp.name
}

func main() {

	var myInt1 int = 222
	fmt.Println(myInt1)
	var myIntParameter1 myInt = 111
	fmt.Println(myIntParameter1.toInt())
	fmt.Println(myIntParameter1.toString())

	var ss anotherInt = "sss"
	fmt.Printf("ss type = %T, value=%s \n", ss, ss)

	var aChinesePerson *ChinesePerson = new(ChinesePerson)
	aChinesePerson.name = "abc"
	fmt.Println(aChinesePerson.sayName())
	fmt.Println(aChinesePerson.sayChineseName())

	var myStringList1 = make(MyStringList, 5)
	fmt.Println(myStringList1.toString())

}

type MyStringList []string

func (p MyStringList) toString() string {
	return "ssss" + strconv.Itoa(len(p)) + "SSS"
}
