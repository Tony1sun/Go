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
	age  int
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
		age:  age,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.name)
}

// 使用值接收者：传拷贝进去
func (p person) guonian() {
	p.age++
}

// 指针接收者：传内存地址进去
// 使用场景 1.需要修改接收者的值  2.接收者是拷贝代价比较大的大对象  3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者
func (p *person) zhengguonian() {
	p.age++
}
func (p *person) dream() {
	fmt.Println("躺平")
}
func main() {
	//d1 := newDog("今天")
	//d1.wang()
	p1 := newPerson("元帅", 18)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.zhengguonian()
	fmt.Println(p1.age)
	p1.dream()
}
