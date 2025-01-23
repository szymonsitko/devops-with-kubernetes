package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var pongCount = 0

var PORT = os.Getenv("PORT")

func main() {
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Printf(">>> Starting web server at port: %s\n", PORT)
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[DEBUG] Incoming request: %s", r.RequestURI)
	}))
	r.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		pongCount += 1
		fmt.Fprintf(w, "Pong %d", pongCount)
	})

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
