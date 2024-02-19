package datastructs

import "fmt"

// Count search depth
var count int

type BinaryNode struct {
	key   int
	left  *BinaryNode
	right *BinaryNode
}

func NewBinarySearchTree() *BinaryNode {
	return &BinaryNode{}
}

func (b *BinaryNode) Insert(value int) {
	if value < b.key {
		if b.left == nil {
			b.left = &BinaryNode{key: value}
		} else {
			b.left.Insert(value)
		}
	} else if value > b.key {
		if b.right == nil {
			b.right = &BinaryNode{key: value}
		} else {
			b.right.Insert(value)
		}
	} else {
		b = &BinaryNode{key: value}
		return
	}
}

func (b *BinaryNode) Search(value int) bool {
	count++
	if b == nil {
		fmt.Println(count)
		count = 0
		return false
	}

	if b.key < value {
		return b.right.Search(value)
	} else if b.key > value {
		return b.left.Search(value)
	}

	fmt.Println(count)
	count = 0
	return true
}

func (b BinaryNode) String() string {
	finalString := fmt.Sprintf("%d %d", b.key, b.right.key)
	return finalString
}
