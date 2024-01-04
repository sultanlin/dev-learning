package data

import (
	"fmt"
)

type Duration float32

type Course struct {
	Id         int
	Name       string
	Slug       string
	Legacy     bool
	Duration   Duration
	Instructor Instructor
}

func (c Course) SignUp() bool {
	return true
}

func (cours Course) String() string {
	return fmt.Sprintf("--- %v --- (%v)\n", cours.Name, cours.Instructor.FirstName)
}
