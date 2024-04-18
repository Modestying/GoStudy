package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sse", handleSSE)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Send an initial "ping" event to establish the connection.
	fmt.Fprintf(w, "data: Initial ping\n\n")

	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Fprintf(w, "data: Time update: %s\n\n", time.Now().Format(time.RFC3339))
			flusher, ok := w.(http.Flusher)
			if !ok {
				log.Println("HTTP response writer does not implement http.Flusher")
				return
			}
			flusher.Flush()
		case <-r.Context().Done():
			log.Println("Client disconnected")
			return
		}
	}
}
