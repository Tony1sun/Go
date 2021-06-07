package main

import (
	"fmt"
	"strings"
)

//闭包

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("fda423ff"))
}
