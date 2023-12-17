package main

import "fmt"

func main() {
	fmt.Print("Hello")
	fmt.Println(" World")
	printData()

	age := 24
	birthday(&age)
	fmt.Println(age)
}

func birthday(age *int) {
	*age++
}
