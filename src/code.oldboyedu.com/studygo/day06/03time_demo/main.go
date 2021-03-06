package main

import (
	"fmt"
	"time"
)

// 时间
func f1() {
	now := time.Now()
	//fmt.Println(now)
	//fmt.Println(now.Year())
	//fmt.Println(now.Month())
	//fmt.Println(now.Day())
	//fmt.Println(now.Minute())
	//fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()
	ret := time.Unix(1623918926, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	// 时间间隔
	fmt.Println(time.Second)
	// now + 24小时
	fmt.Println(now.Add(24 * time.Hour))
	// Sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02 15:04:05", "2020-01-01 12:25:00")
	if err != nil {
		fmt.Printf("parse time failed, err:%V\n", err)
		return
	}
	d := now.Sub(nextYear)
	fmt.Println(d)
	fmt.Println("-------------")
	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer{
	//	fmt.Println(t)
	//}
	// 格式化时间 把语言中的时间对象转换成字符串类型的时间
	// 2021-06-17
	fmt.Println(now.Format("2006-01-02"))
	// 2021/06/17 16:49:00
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2021/06/17 16:49:00
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))
	// 2021/06/17 16:49:00.342
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	// 按照对应的格式解析字符串类型的时间
	//timeObj, err := time.Parse("2006-01-02", "2017-08-03")
	//if err != nil {
	//	fmt.Println("parse time failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println(timeObj)
	//fmt.Println(timeObj.Unix())

	// Sleep
	n := 5
	fmt.Println("开始sleep了")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5秒钟过去了")
}

//时区
func f2() {
	now := time.Now() // 本地的时间
	fmt.Println(now)
	// 明天的这个时间
	time.Parse("2006-01-02 15:04:05", "2020-06-17 14:51:50")
	// 按照东八区的时区和格式取解析一个字符串格式的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err:%v\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-06-17 14:51:50", loc)
	if err != nil {
		fmt.Printf("load loc failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}
func main() {
	//f1()
	f2()
}
