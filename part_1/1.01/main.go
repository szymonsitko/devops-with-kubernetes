package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var PORT = os.Getenv("PORT")

var (
	currentTime string
	hash        string
)

func http_server() {
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Printf(">>> Starting web server at port: %s\n", PORT)
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[DEBUG] Incoming request: %s", r.RequestURI)
	}))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, hash)
	})

	http.ListenAndServe(":"+PORT, r)
}

func main() {
	go http_server()

	for {
		currentTime = time.Now().UTC().Format(time.RFC3339)
		hash = fmt.Sprintf("%s: %s\n", currentTime, uuid.New())
		fmt.Print(hash)
		time.Sleep(5 * time.Second)
	}
}
