package main

import (
	"time"
)
import . "log"

func main() {
	//获取当前时间  时间分为时间点Time和时间段Duration
	//包括年月日时分秒，毫秒，市区，还有一堆看不懂的东西
	var now1 = time.Now()
	Println(now1)

	//时间点的格式化 谁能想到2006-01-02 15:04:05就是时间格式的规则表达式写法呢，原来是从1234567的顺序排列下来的
	Println(now1.Format("2006-01-02 15:04:05"))
	//输出的就是一个yyyy-MM-dd HH:mm:ss格式的字符串

	//新建一个日期+时间
	var t1 = time.Date(2021, 12, 25, 13, 22, 59, 231, time.Local)
	Println(t1)
	//获取时间戳
	Println(t1.Unix()) //秒时间戳
	// time.Unix的第二个参数传递0或10结果一样，因为都不大于1e9
	var unixInt64ToTime = time.Unix(t1.Unix(), 0)
	Println("###:", unixInt64ToTime)
	Println(t1.UnixNano())  //纳秒时间戳
	Println(t1.UnixMicro()) //微秒时间戳
	Println(t1.UnixMilli()) //毫秒时间戳

	//字符串转时间
	var t2, er1 = time.Parse("2006-01-02 15:04:05", "2098-07-06 17:04:03")
	Println(t2, er1)

	//提到时间，必须要有确定时区
	var loc1, err1 = time.LoadLocation("")
	Println(loc1, err1)
	loc1, err1 = time.LoadLocation("Local")
	Println(loc1, err1)

	//时间段，就是一段时间。是两个时间点之间的时间
	var duration1, _ = time.ParseDuration("3s")
	Println(duration1)                //默认调用的是string格式
	Println(duration1.Milliseconds()) //毫秒单位
	var start = time.Now().UnixMilli()
	time.After(duration1) //after不阻塞
	Println("time.After ", time.Now().UnixMilli()-start, "ms")
	//非阻塞延迟处理，当延迟的时间段到时间后，就会执行函数中的内容
	time.AfterFunc(duration1, func() {
		Println("time.AfterFunc(duration1, func go here")
	})
	start = time.Now().UnixMilli()
	time.Sleep(time.Duration(4) * time.Second) //sleep是阻塞的
	Println("time.Sleep ", time.Now().UnixMilli()-start, "ms")

	var BASE_TIME = time.Date(2021, 12, 25, 13, 11, 10, 1, time.Local)
	Println(time.Since(BASE_TIME)) //从basetime开始到现在，一共经历的多少小时多少分钟多少秒
	Println(time.Until(BASE_TIME)) //从现在开始直到basetime结束，一共经历的多少小时多少分钟多少秒

}
