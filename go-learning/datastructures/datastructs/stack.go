package datastructs

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	removedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return removedItem
}

func (s Stack) String() string {
	finalString := ""
	for _, item := range s.items {
		finalString += fmt.Sprintf("%d ", item)
	}
	finalString += fmt.Sprintf("\n")
	return finalString
}
