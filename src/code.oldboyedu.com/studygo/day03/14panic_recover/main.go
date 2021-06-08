package main

import "fmt"

func funcA() {
	fmt.Println("func_A")
}

func funcB() {
	fmt.Println("func_b")
}

func funcC() {
	fmt.Println("func_C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
