package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	w.WriteHeader(http.StatusOK)

	count := 1

	for {
		fmt.Fprintf(w, "data: %d\n\n", count)
		w.(http.Flusher).Flush()
		count += 1
		time.Sleep(2 * time.Second)
	}
}
