package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("Calculator 1.0")
	fmt.Println("Which operation do you want?")

	fmt.Scanf("%s", &operation)
	fmt.Println("Enter the first number")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter the second number")
	fmt.Scanf("%d", &num2)

	fmt.Printf("The answer to the %s of %d and %d is: ", operation, num1, num2)
	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "subtract":
		fmt.Println(num1 - num2)
	case "multiply":
		fmt.Println(num1 * num2)
	case "divide":
		fmt.Println(num1 / num2)
	}

}
