package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
	w.Write([]byte("Hello World"))
}

func main() {
	fmt.Println("This is a test")
	// server := http.NewServeMux()
	server := http.NewServeMux()

	// server.HandleFunc("/", hello)
	// currentProj := http.FileServer(http.Dir("./first"))
	// server.Handle("/", currentProj)
	// bookProj := http.FileServer(http.Dir("./bookexample"))
	// server.Handle("/", bookProj)
	// server.Handle("/book", bookProj)
	finalProj := http.FileServer(http.Dir("./experiments/"))
	server.Handle("/", finalProj)

	err := http.ListenAndServe(":3333", server)

	fmt.Println(err)
}
