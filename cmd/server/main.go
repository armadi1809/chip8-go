package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting file server on 8080")
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./web-assets")))
	if err != nil {
		log.Fatalf("error occurred, server shutting down %v\n", err)
	}
}
