package main

import (
	"fmt"
	"net/http"
)

// func quiz(w http.ResponseWriter, r *http.Request) {
//
// }

func main() {
	server := http.NewServeMux()

	quizFS := http.FileServer(http.Dir("./firstsetup"))
	// doggoFS := http.FileServer(http.Dir("./doggoquiz"))
	server.Handle("/", quizFS)
	// server.Handle("/", doggoFS)
	// server.HandleFunc("/", quiz)

	err := http.ListenAndServe(":3333", server)

	fmt.Println(err)
}
