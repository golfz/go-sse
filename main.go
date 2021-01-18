package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handlerRequest()
}

func handlerRequest() {
	http.HandleFunc("/sse", sseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
