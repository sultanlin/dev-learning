package main

import (
	"fmt"

	"sultan.com/go/server/data"
)

func main() {
	max := data.Instructor{Id: 30, LastName: "Linjawi"}
	max.FirstName = "Sultan"

	fmt.Println(max.Print())

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: max}

	// fmt.Printf(goCourse)
	// fmt.Printf("%v", goCourse)
	// fmt.Printf(goCourse.String())

	// swiftWS := data.Workshop{
	// 	Course: data.Course{Name: "Swift for iOS", Instructor: max},
	// 	Date:   time.Now(),
	// }

	swiftWS := data.NewWorkshop("Swift for iOS", max)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}

	// fmt.Printf("%v", swiftWS)
	// fmt.Printf("%v", swiftWS.String())
}
