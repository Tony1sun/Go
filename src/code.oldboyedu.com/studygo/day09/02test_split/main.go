package main

import (
	"code.oldboyedu.com/studygo/day09/split_string"
	"fmt"
)

func main() {
	ret := split_string.Split("bbdbfds", "b")
	fmt.Printf("%#v\n", ret)

	ret2 := split_string.Split("b1dbsb", "b")
	fmt.Printf("%#v\n", ret2)

	ret3 := split_string.Split("bfdbbfds", "b")
	fmt.Printf("%#v\n", ret3)
}
