package mycontainer

import "fmt"

type SNode struct {
	Value int
	Next  *SNode
}

var Stack = new(SNode)
var SSize int

func PushS(payload int) bool {
	if SSize == 0 {
		Stack = &SNode{payload, nil}
		SSize++
		return true
	}

	temp := &SNode{payload, Stack}
	Stack = temp
	SSize++
	return true
}

func PopS(s *SNode) (int, bool) {
	if SSize == 0 {
		return 0, false
	}

	if SSize == 1 {
		Stack = nil
		SSize--
		return s.Value, true
	}

	Stack = Stack.Next
	SSize--
	return s.Value, true
}

func TraverseS(s *SNode) {
	if s == nil {
		fmt.Println("Empty Stack!")
	}
	for s != nil {
		fmt.Print(s.Value, "->")
		s = s.Next
	}
	fmt.Println()
}
