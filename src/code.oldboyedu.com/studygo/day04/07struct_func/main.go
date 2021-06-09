package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

// 当结构体比较大时尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}
func main() {
	p1 := newPerson("今天", 30)
	p2 := newPerson("明天", 18)
	fmt.Println(p1, p2)
	d1 := newDog("后天")
	fmt.Println(d1)
}
