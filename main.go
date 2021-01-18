package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CountNumber struct {
	Number int       `json:"number"`
	Time   time.Time `json:"time"`
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.WriteHeader(http.StatusOK)

	count := 1

	for {
		d := CountNumber{
			Number: count,
			Time:   time.Now(),
		}

		s, _ := json.Marshal(d)

		fmt.Fprintf(w, "data: %s\n\n", s)
		w.(http.Flusher).Flush()
		count += 1
		time.Sleep(2 * time.Second)
	}
}

func handlerRequest() {
	http.HandleFunc("/sse", sseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handlerRequest()
}
