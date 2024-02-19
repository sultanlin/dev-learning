package datastructs

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l LinkedList) String() string {
	var finalPrint string
	for toPrint := l.Head; l.Length != 0; l.Length-- {
		finalPrint += fmt.Sprintf("%d ", toPrint.Data)
		toPrint = toPrint.Next
	}
	fmt.Println()
	return finalPrint
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		fmt.Println("Empty list")
		return
	}
	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.Length--
		fmt.Println("Value found in header")
		return
	}
	previousToDelete := l.Head
	for previousToDelete.Next.Data != value {
		if previousToDelete.Next.Next == nil {
			fmt.Println("Value not found")
			return
		}
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.Length--
}
