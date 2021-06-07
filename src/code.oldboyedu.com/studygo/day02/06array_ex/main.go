package main

import "fmt"

func main() {
	//求数组所有元素的和
	var c1 = [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v1 := range c1 {
		sum += v1
	}
	fmt.Println(sum)

	//找出和为8的两个元素的下标分别为 (0,3)(1,2)
	//定义两个for 循环, 外层的从第一个开始遍历
	//内层的for循环从外层后面的那个开始找

	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			if c1[i]+c1[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
