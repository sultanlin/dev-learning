package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Hello("Sultan").Render(r.Context(), w)
	})
	http.ListenAndServe(":3333", nil)
}
