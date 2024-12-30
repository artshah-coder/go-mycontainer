package mycontainer

import (
	"fmt"
)

type HTNode struct {
	Value int
	Next  *HTNode
}

type HashTable struct {
	Table map[int]*HTNode
	Size  int
}

func hashFunction(value, size int) int {
	return value % size
}

func InsertHT(hash *HashTable, value int) int {
	index := value % hash.Size
	HTNode := HTNode{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &HTNode
	return index
}

func TraverseHT(hash *HashTable) {
	for _, HTNode := range hash.Table {
		for HTNode != nil {
			fmt.Print(HTNode.Value, " ->")
			HTNode = HTNode.Next
		}
		fmt.Println()
	}
}

func LookupHT(hash *HashTable, value int) bool {
	index := hashFunction(value, hash.Size)
	for HTNode := hash.Table[index]; HTNode != nil; HTNode = HTNode.Next {
		if HTNode.Value == value {
			return true
		}
	}
	return false
}
