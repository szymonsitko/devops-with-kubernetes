package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT = os.Getenv("PORT")

func main() {
	if PORT == "" {
		log.Fatal("PORT variable must be set! Terminating...")
	}
	fmt.Printf(">>> Serving app on PORT: %s\n", PORT)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("content-type", "application/html")
		fmt.Fprintf(w, "<h1>Hello, you've requested: %s</h1>", r.URL.Path)
	})

	http.ListenAndServe(":"+PORT, nil)
}
