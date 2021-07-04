package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverselList(head *ListNode) *ListNode {
	// 前一个节点
	var pre *ListNode
	//cur := head
	for head != nil {
		cur := head.Next // 当前节点指向下一个节点
		head.Next = pre  // 将当前节点指向前一个节点
		pre = head
		head = cur
	}
	return pre
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Printf("%#v\n", head)
	ret := reverselList(head)
	fmt.Printf("%#v\n", ret)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
}
