package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Beginning server")
	server := http.NewServeMux()
	formsFS := http.FileServer(http.Dir("./pages"))
	server.Handle("/", formsFS)
	err := http.ListenAndServe(":3333", server)

	fmt.Println(err)
}
