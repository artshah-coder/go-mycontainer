package mycontainer

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func TraverseBT(tree *Tree) {
	if tree == nil {
		return
	}

	TraverseBT(tree.Left)
	fmt.Print(tree.Value, "	")
	TraverseBT(tree.Right)
}

func CreateBT(n int) *Tree {
	var tree *Tree
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 2*n; i++ {
		tree = InsertBT(tree, rand.Intn(2*n))
	}
	return tree
}

func InsertBT(tree *Tree, n int) *Tree {
	if tree == nil {
		return &Tree{nil, n, nil}
	}
	if tree.Value == n {
		return tree
	}
	if n < tree.Value {
		tree.Left = InsertBT(tree.Left, n)
		return tree
	}
	tree.Right = InsertBT(tree.Right, n)
	return tree
}

func LookupBT(tree *Tree, n int) bool {
	if tree == nil {
		return false
	}
	if tree.Value == n {
		return true
	}
	if n < tree.Value {
		return LookupBT(tree.Left, n)
	}
	return LookupBT(tree.Right, n)
}
