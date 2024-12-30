package mycontainer

import (
	"fmt"
)

type LNode struct {
	Value int
	Next  *LNode
}

var RootL = new(LNode)

func AddLNode(list *LNode, payload int) int {
	if list == nil {
		list = &LNode{Value: payload, Next: nil}
		RootL = list
		return 0
	}

	if list.Value == payload {
		fmt.Println("Element alrady exists:", payload)
		return -1
	}

	if list.Next == nil {
		list.Next = &LNode{Value: payload, Next: nil}
		return 0
	}

	if payload < list.Next.Value && payload > list.Value {
		tmp := list.Next
		list.Next = &LNode{Value: payload, Next: tmp}
		return 0
	}

	return AddLNode(list.Next, payload)
}

func DelLNode(list *LNode, payload int) int {
	if list == nil {
		fmt.Println("Element does not exist:", payload)
		return -1
	}

	if list.Next != nil && list.Next.Value == payload {
		list.Next = list.Next.Next
		return 0
	} else if list.Next == nil {
		if list.Value == payload {
			list = nil
			RootL = list
			return 0
		}
	}

	return DelLNode(list.Next, payload)
}

func TraverseL(list *LNode) {
	if list == nil {
		fmt.Println("List is empty")
		return
	}
	for l := list; l != nil; l = l.Next {
		fmt.Print(l.Value, "->")
	}
	fmt.Println()
}

func LookupLNode(list *LNode, payload int) bool {
	if list == nil {
		return false
	}

	if list.Value == payload {
		return true
	}

	return LookupLNode(list.Next, payload)
}

func SizeL(list *LNode) int {
	size := 0
	for l := list; l != nil; l = l.Next {
		size++
	}
	return size
}
