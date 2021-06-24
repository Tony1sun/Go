package main

import (
	"fmt"
	"reflect"
)

//func reflectType(x interface{}) {
//	v := reflect.TypeOf(x)
//	fmt.Printf("type:%v\n", v)
//}
//
//
//func main() {
//	var a float32 =  3.14
//	reflectType(a)
//	var b int64 = 100
//	reflectType(b)
//}

// type name和type kind

//type myInt int64
//
//func reflectType(x interface{}) {
//	t := reflect.TypeOf(x)
//	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
//}
//
//
//func main() {
//	var a *float32 //指针
//	var b myInt //自定义类型
//	var c rune //类型别名
//	reflectType(a)
//	reflectType(b)
//	reflectType(c)
//
//	type person struct {
//		name string
//		age int
//	}
//
//	type book struct {
//		title string
//	}
//
//	var d = person{
//		name: "今天",
//		age: 18,
//	}
//	var e = book{
//		title: "下雨",
//	}
//	reflectType(d)
//	reflectType(e)
//}

//通过反射获取值
//func reflectValue(x interface{}) {
//	v := reflect.ValueOf(x)
//	k := v.Kind()
//	switch k {
//	case reflect.Int64:
//		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
//		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
//	case reflect.Float32:
//		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
//		fmt.Printf("type is int32, value is %f\n", float32(v.Float()))
//	case reflect.Float64:
//		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
//		fmt.Printf("type is int64, value is %f\n", float64(v.Float()))
//	}
//}
//
//
//func main() {
//	var a float32 = 3.14
//	var b int64 = 100
//	reflectValue(a)
//	reflectValue(b)
//	// 将int类型的原始值转换为reflect.Value类型
//	c := reflect.ValueOf(10)
//	fmt.Printf("type c :%T\n", c)
//}

//通过反射设置变量的值

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本, reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100
	//reflectSetValue1(a)  //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
