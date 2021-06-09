package main

import "fmt"

// 方法
// Go语言中如果标识符首字母是大写的，就表示对外部可见
// dog 这是一个狗的结构体

type dog struct {
	name string
}


type person struct {
	name string
	age int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age: age,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.name)
}

// 使用值接收者：传拷贝进去
func (p person) guonian {
	p.age++
}
func main() {
	d1 := newDog("今天")
	d1.wang()
}

