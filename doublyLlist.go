package mycontainer

import (
	"fmt"
)

type DLNode struct {
	Value    int
	Previous *DLNode
	Next     *DLNode
}

var RootDL = new(DLNode)

func AddDLNode(list *DLNode, payload int) int {
	if list == nil {
		list = &DLNode{payload, nil, nil}
		RootDL = list
		return 0
	}

	if list.Value == payload {
		fmt.Println("Element already exists:", payload)
		return -1
	}

	if list.Value < payload && list.Next != nil && list.Next.Value > payload {
		temp := &DLNode{payload, list, list.Next}
		list.Next.Previous = temp
		list.Next = temp
		return 0
	}

	if list.Value >= payload {
		temp := &DLNode{payload, list.Previous, list}
		if list.Previous != nil {
			list.Previous.Next = temp
		}
		list.Previous = temp
		RootDL = temp
		return 0
	}

	if list.Next == nil {
		temp := list
		list.Next = &DLNode{payload, temp, nil}
		return 0
	}

	return AddDLNode(list.Next, payload)
}

func DelDLNode(list *DLNode, payload int) int {
	if list == nil {
		fmt.Println("Element does not exists:", payload)
		return -1
	}

	if list.Value == payload {
		switch {
		case list.Next != nil && list.Previous != nil:
			list.Previous.Next = list.Next
			list.Next.Previous = list.Previous
		case list.Previous == nil && list.Next == nil:
			list = nil
			RootDL = list
		case list.Previous == nil:
			list.Next.Previous = nil
		case list.Next == nil:
			list.Previous.Next = nil
		}
		return 0
	}

	return DelDLNode(list.Next, payload)
}

func TraverseDL(list *DLNode) {
	if list == nil {
		fmt.Println("Empty list")
		return
	}

	for l := list; l != nil; l = l.Next {
		fmt.Print(l.Value, "->")
	}
	fmt.Println()
}

func ReverseDL(list *DLNode) {
	if list == nil {
		fmt.Println("Empty list")
		return
	}

	var temp *DLNode
	for ; list != nil; list = list.Next {
		temp = list
	}
	for ; temp != nil; temp = temp.Previous {
		fmt.Print(temp.Value, "->")
	}
	fmt.Println()
}

func SizeDL(list *DLNode) int {
	if list == nil {
		return 0
	}

	size := 0
	for ; list != nil; list = list.Next {
		size++
	}
	return size
}

func LookupDLNode(list *DLNode, payload int) bool {
	if list == nil {
		return false
	}

	if list.Value == payload {
		return true
	}

	return LookupDLNode(list.Next, payload)
}
