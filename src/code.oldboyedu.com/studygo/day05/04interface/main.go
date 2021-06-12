package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车，都能跑

type car interface {
	run()
}

type falali struct {
	brand string
}

func (f falali) run() {
	fmt.Printf("%s速度10迈", f.brand)
}

type baoshijie struct {
	brand string
}

func (f baoshijie) run() {
	fmt.Printf("%s速度10迈", f.brand)
}

func driver(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	driver(f1)
	driver(b1)

}
