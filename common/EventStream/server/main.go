package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func stream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	id := 0
	for {
		select {
		case <-ctx.Done():
			log.Printf("Connection closed | IP: %s", r.RemoteAddr)
			return

		case t := <-ticker.C:
			id++
			fmt.Fprintf(w, "id: %d\n", id)
			fmt.Fprintf(w, "event: message\n")
			fmt.Fprintf(w, "data: %s\n\n", t.Format(time.RFC3339))
			flusher.Flush()
		}
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"status":"ok","endpoint":"/stream"}`)); err != nil {
		log.Printf("Write error: %v", err)
	}
}

func main() {
	http.HandleFunc("/stream", stream)
	http.HandleFunc("/health", health)
	log.Println("SSE server started | visit: http://localhost:8080/stream")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
