package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (instr Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)", instr.LastName, instr.FirstName, instr.Score)
}
