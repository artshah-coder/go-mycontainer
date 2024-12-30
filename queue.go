package mycontainer

import "fmt"

type QNode struct {
	Value int
	Next  *QNode
}

var Queue = new(QNode)
var QSize int

func PushQ(q *QNode, payload int) bool {
	if q == nil {
		q = &QNode{payload, nil}
		Queue = q
		QSize++
		return true
	}

	temp := &QNode{payload, q}
	Queue = temp
	QSize++
	return true
}

func PopQ(q *QNode) (int, bool) {
	if q == nil {
		return 0, false
	}

	if QSize == 1 {
		Queue = nil
		QSize--
		return q.Value, true
	}

	var temp *QNode
	for q.Next != nil {
		temp = q
		q = q.Next
	}

	v := temp.Next.Value
	temp.Next = nil
	QSize--

	return v, true
}

func TraverseQ(q *QNode) {
	for q != nil {
		fmt.Print(q.Value, "->")
		q = q.Next
	}
	fmt.Println()
}
