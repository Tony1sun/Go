package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverselList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next // 把下一个节点缓存起来
		cur.Next = pre
		pre = cur
		cur = tmp
		//pre, cur, cur.Next = cur, cur.Next, pre
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
