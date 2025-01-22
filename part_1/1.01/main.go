package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	for {
		currentTime := time.Now().UTC().Format(time.RFC3339)
		fmt.Printf("%s: %s\n", currentTime, uuid.New())
		time.Sleep(5 * time.Second)
	}
}
