package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi, This is b29c7aa!\n")
	})
	http.ListenAndServe(":8080", nil)
}
