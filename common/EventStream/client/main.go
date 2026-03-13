package main

import (
	"bufio"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "http://localhost:8080/stream"
	log.Printf("Connecting to SSE: %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("request create error: %v", err)
	}
	req.Header.Set("Accept", "text/event-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("http status: %d", resp.StatusCode)
	}

	log.Println("Connected. Receiving events...")

	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("read error: %v", err)
			return
		}

		line = strings.TrimRight(line, "\r\n")

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, ":") {
			continue
		}

		if strings.HasPrefix(line, "data:") {
			value := strings.TrimSpace(line[len("data:"):])
			log.Printf("data: %s", value)
		}
	}
}
