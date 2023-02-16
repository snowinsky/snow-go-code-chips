package main

import "fmt"

func main() {
	//map make[key]value
	//map的创建 使用make关键字
	var map1 map[int]string = make(map[int]string)
	var map2 map[string]int = make(map[string]int)

	//map的添加元素，类似于二维数组的玩法，key是数组下标，value就是数组存的值
	map1[0] = "zero"
	map1[1] = "one"
	map1[2] = "two"

	map2["three"] = 3
	map2["four"] = 4
	map2["five"] = 5

	//map的删除元素，使用delete关键字
	delete(map1, 1)
	delete(map2, "four")

	//map的遍历元素
	fmt.Println(map1)
	fmt.Println(map2)

	//map的遍历是通过range关键字
	//对map来讲，range遍历的是key, range返回的就是这个map的key，然后通过key来获取到对应的value即可
	//range 迭代map时，返回的是key
	for key := range map1 {
		fmt.Println(key, map1[key])
	}

	//map的是否存在元素，通过map的key直接获取value时同时返回是否存在的bool值，可以先判断是否存在，然后再拿取值。
	var anyValue, isExists = map2["five"]
	fmt.Println(anyValue, isExists)

	//map是不允许null key和null value的，在编译阶段就不会通过
	/*var myMapBean = make(map[int]nil)
	myMapBean[0] = MapBean{0}
	myMapBean[1] = MapBean{1}
	myMapBean[-1] = nil*/

}
