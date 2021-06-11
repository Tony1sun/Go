package main

import (
	"encoding/json"
	"fmt"
)

// 复习结构体
//有名字的结构体
type tmp struct {
	x int
	y int
}

type person struct {
	name  string
	age   int
	hobby string
}

// 构造(结构体变量的)函数
func newPerson(n string, i int, h string) person {
	return person{
		name:  n,
		age:   i,
		hobby: h,
	}
}

// 方法
// 接收者使用对应类型的首字母小写
// 指定了接收者之后，只有接收者这个类型的变量才能调用这个方法
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是%s.\n", p.name, str)
}

//func (p person) guonian() {
//	p.age++ // 此处的p是p1的副本，改的是副本
//}

//指针接收者
// 1.需要修改结构体变量的值要使用指针接收者
// 2.结构体本身比较多，拷贝的内存开销比较大时也要使用指针接收者
// 3.保持一致性；如果有一个方法使用了指针接收者，其他的方法为了统一也要使用指针接收者
func (p *person) guonian() {
	p.age++ // 此处的p是p1的副本，改的是副本
}

func main() {
	var b = tmp{
		10,
		20,
	}
	fmt.Println(b)

	// 匿名结构体
	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	var x int
	y := int8(10)
	fmt.Println(x, y)

	var p1 person //结构体实例化
	p1.name = "今天"
	p1.age = 18

	p2 := person{"保德路", 18, "爬山"} // 结构体实例化
	p3 := person{"保德路", 18, "爬山"}
	fmt.Println(p1, p2)

	// 调用构造函数生成person类型变量
	p4 := newPerson("今天", 30, "游泳")
	fmt.Println(p1, p2, p3, p4)

	p1.dream("泡妹子")
	p2.dream("还是泡妹子")

	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)

	// 结构体嵌套
	type addr struct {
		province string
		cit      string
	}
	type student struct {
		name    string
		address addr // 嵌套别的结构体
	}

	// 序列化
	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
	po1 := point{100, 200}
	b1, err := json.Marshal(po1)
	// 如果出错了
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Println(string(b1))

	// 反序列化：由字符串 --> Go中结构体变量
	str1 := `{"x":10,"y":20}`
	var po2 point // 造一个结构体变量，准备接受反序列化的值
	err = json.Unmarshal([]byte(str1), &po2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	}
	fmt.Println(po2)
}
