package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Welcome to the Go image transformation service!")
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
