package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello this is a demo backend from Golang")
	})
	fmt.Println("server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
