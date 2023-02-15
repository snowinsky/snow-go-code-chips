package main

import (
	"fmt"
	"sort"
)

// go里边不固定的数组叫切片，是不是list被申请专利了呀，就天天新造概念玩
func main() {
	//切片的新建
	var intList1 []int
	fmt.Println(intList1)
	var intList2 []int = []int{}
	fmt.Println(intList2)
	var intList3 []int = []int{3, 1, 2}
	fmt.Println(intList3)
	var intList4 []int = make([]int, 0)
	fmt.Println(intList4)
	fmt.Println(append(intList4, 1, 2, 3, 4, 5))

	//切片的截取 subList
	intListSearch := make([]int, 0)
	intListSearch = append(intListSearch, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(intListSearch, intListSearch[0], intListSearch, intListSearch[3])

	//切片的合并
	newIntListSearch := make([]int, 0)
	newIntListSearch = append(newIntListSearch, intListSearch[0:3]...)
	fmt.Println(newIntListSearch)
	newIntListSearch = append(newIntListSearch, intListSearch[7:]...)
	fmt.Println(newIntListSearch)

	//切片的删除元素
	intListForDelete := make([]int, 0)
	fmt.Println(append(intListForDelete, 1, 2, 3, 4, 5))
	fmt.Println(intListForDelete) //执行到这里，list中根本没有元素
	intListForDelete = append(intListForDelete, 1, 2, 3, 4, 5)
	fmt.Println(intListForDelete) //执行到这里，list才有元素

	//切片的新加元素
	stringList1 := make([]string, 0)
	fmt.Println(append(stringList1, "a", "b"))
	stringList1 = append(stringList1, "a", "b", "a")
	fmt.Println(stringList1)

	//切片的元素排序
	//常见的元素的排序都给了封装
	sort.Strings(stringList1)
	fmt.Println(stringList1)

	sort.Ints(intListSearch)
	fmt.Println(intListSearch)

	//如果是一个特殊的对象要排序，那这个list必须实现一个接口中的三个方法，好麻烦
	var instanceStringList = StringList([]string{"b", "a", "b"})
	sort.Sort(instanceStringList)
	fmt.Println(instanceStringList)

}

type StringList []string

func (p StringList) Len() int {
	return len(p)
}
func (p StringList) Less(i, j int) bool {
	return p[i] > p[j]
}
func (p StringList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
