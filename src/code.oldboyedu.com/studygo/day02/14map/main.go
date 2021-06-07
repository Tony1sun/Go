package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10) //要估算好该map容量，避免在程序运行期间再动态扩容
	m1["今天"] = 65
	m1["明天"] = 66
	fmt.Println(m1)
	fmt.Println(m1["今天"])

	value, ok := m1["后天"]
	if !ok {
		fmt.Println("查无此KEY")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除
	delete(m1, "今天")
	fmt.Println(m1)
	delete(m1, "沙河") //删除不存在的key
}
