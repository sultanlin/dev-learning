package main

import "sultan.com/go/server/data"

func main() {
	max := data.Instructor{Id: 30, LastName: "Linjawi"}
	max.FirstName = "Sultan"

	print(max.Print())
}
