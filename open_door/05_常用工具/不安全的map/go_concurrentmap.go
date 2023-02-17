package main

import (
	"fmt"
	"sync"
)

// go中的map并不是线程安全的，高并发情况下会出问题，特引入sync.Map
func main() {
	concurrentMap := sync.Map{}
	//插入
	concurrentMap.Store(1, "one")
	concurrentMap.Store(2, "two")
	fmt.Println(concurrentMap.Load(2))
	//更新
	concurrentMap.Store(2, "二")
	fmt.Println(concurrentMap.Load(2))
	//删除
	concurrentMap.Delete(2)
	fmt.Println(concurrentMap.Load(2))
	//遍历
	//key存在就返回value，如果不存在就插入value
	v1, hasV1 := concurrentMap.LoadOrStore(1, "一")
	fmt.Println(v1, hasV1)
	fmt.Println(concurrentMap.Load(1))
	v3, hasV3 := concurrentMap.LoadOrStore(3, "three")
	fmt.Println(v3, hasV3)
	fmt.Println(concurrentMap.Load(3))

	concurrentMap.Range(func(key interface{}, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	//key存在就返回value，并删除；如果不存在就返回null
	v1, hasV1 = concurrentMap.LoadAndDelete(1)
	fmt.Println(v1, hasV1)
	fmt.Println(concurrentMap.Load(1))

	v4, hasV4 := concurrentMap.LoadAndDelete(4)
	fmt.Println(v4, hasV4)

	concurrentMap.Range(func(key interface{}, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
