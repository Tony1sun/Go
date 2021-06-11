package main

import "fmt"
import "student_mgr"

// 学生管理系统
var smr studentMgr //声明一个全局变量：学生管理对象smr

func showMenu() {
	fmt.Println("welcome sms!")
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[string]student, 100),
	}
	for {
		showMenu()
		// 等待用户输入
		fmt.Print("请输入序号:")
		var choice int
		fmt.Scanln()
		fmt.Println("你输入的是:", choice)
		switch choice {
		case 1:
			smr.show
		}

	}
}
