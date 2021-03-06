package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动!")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡冻!")
}
func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s...\n", food)
}
func main() {
	var a1 animal

	bc := cat{ //定义一个cat类型的变量bc
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	a1.eat("大鱼")
	fmt.Printf("%T\n", a1)
	fmt.Println(a1)

	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	a1.eat("饲料")
	fmt.Printf("%T\n", a1)
	fmt.Println(a1)
}
