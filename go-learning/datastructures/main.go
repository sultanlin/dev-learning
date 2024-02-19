package main

import (
	"fmt"

	"github.com/sultanlinjawi/linkedlist/datastructs"
)

func linkedlists() {
	myList := datastructs.LinkedList{}

	node1 := &datastructs.Node{Data: 4}
	node2 := &datastructs.Node{Data: 20}
	node3 := &datastructs.Node{Data: 13}

	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)

	myList.Prepend(&datastructs.Node{Data: 11})
	myList.Prepend(&datastructs.Node{Data: 25})
	myList.Prepend(&datastructs.Node{Data: 18})
	myList.Prepend(&datastructs.Node{Data: 43})
	fmt.Println(myList)

	myList.DeleteWithValue(13)
	myList.DeleteWithValue(10000)
	myList.DeleteWithValue(43)

	fmt.Println(myList)
}

func stacks() {
	fmt.Println("Stacks start")

	myStack := datastructs.Stack{}

	myStack.Push(5)
	myStack.Push(10)
	myStack.Push(26)
	myStack.Push(2)
	fmt.Println(myStack)

	checkReturned := myStack.Pop()
	fmt.Println(checkReturned)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}

func queues() {
	fmt.Println("\nQueues start")

	myQueue := datastructs.Queue{}

	myQueue.Enqueue(5)
	myQueue.Enqueue(10)
	myQueue.Enqueue(26)
	myQueue.Enqueue(2)
	fmt.Println(myQueue)

	checkReturned := myQueue.Dequeue()
	fmt.Println(checkReturned)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}

func binarySearchTree() {
	fmt.Println("Binary search tree--------")
	tree := datastructs.NewBinarySearchTree()
	// tree := datastructs.BinaryNode{}

	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(17)
	tree.Insert(1)

	fmt.Println(tree)
	fmt.Println(tree.Search(10))
	fmt.Println(tree.Search(1))
	fmt.Println(tree.Search(8))
	fmt.Println(tree.Search(18))
}

func hashTables() {
	// testHashTable := &datastructs.HashTable{}
	testHashTable := datastructs.NewHashTable()
	fmt.Println(testHashTable)
	fmt.Println(datastructs.Hash("RANDY"))
}

func main() {
	// linkedlists()
	// stacks()
	// queues()

	// binarySearchTree()
	hashTables()
}
