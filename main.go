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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintln(w, "Hello")
}
