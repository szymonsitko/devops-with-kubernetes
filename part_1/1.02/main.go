package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT = os.Getenv("PORT")

func response(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("Server started in port %s\n", PORT))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	if PORT == "" {
		log.Fatal("PORT variable must be set, terminating...")
	}

	fmt.Printf("Server started in port %s\n", PORT)

	http.HandleFunc("/headers", headers)
	http.HandleFunc("/response", response)

	http.ListenAndServe(":"+PORT, nil)
}
