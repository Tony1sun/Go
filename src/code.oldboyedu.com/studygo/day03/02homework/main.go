package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1.判断字符串中汉字的数量
	// 难点是判断一个字符是汉字
	s1 := "Hello今天にほんご"
	// 1.依次拿到字符串中的字符
	var count int
	for _, c := range s1 {
		// 2.判断当前字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++

		}
	}
	// 3.把汉字出现的次数累加得到最终结果
	fmt.Println(count)

	// 2.how do you do 单词出现的次数
	s2 := "how do you do"
	// 2.1 把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	// 2.2遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		// 1.如果原来map中不存在w这个key，那么出现次数 = 1
		// 1.如果原来map中存在w这个key，那么出现次数 + 1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	// 2.3累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	// 回文判断
	// 字符串从左往右读和从左往右读是一样的
	// 上海自来水来自海上

	ss := "上海自来水来自海上"
	// 把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		// 上 r[0] r[len(r)-1]
		// 海 r[1] r[len(r)-1-1]
		// 自 r[2] r[len(r)-1-2]
		// c r[i] r[len(r)-1-i]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
