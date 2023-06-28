package main

import (
	"fmt"
	"log"
	"net/http"

	"todo/router"
) 

var token = "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"
func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
