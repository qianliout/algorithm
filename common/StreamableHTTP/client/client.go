package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type clientConfig struct {
	URL     string
	Timeout time.Duration
}

func newClientConfig() clientConfig {
	url := strings.TrimSpace(os.Getenv("MCP_URL"))
	if url == "" {
		url = "http://localhost:9090/mcp"
	}
	return clientConfig{
		URL:     url,
		Timeout: 30 * time.Second,
	}
}

type jsonRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type streamMessage struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   json.RawMessage `json:"error,omitempty"`
}

type clientInputNotification struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

func main() {
	cfg := newClientConfig()
	log.Printf("starting client url=%s", cfg.URL)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	pr, pw := io.Pipe()
	enc := json.NewEncoder(pw)

	go func() {
		defer func() {
			_ = pw.Close()
		}()

		first := jsonRPCRequest{
			JSONRPC: "2.0",
			ID:      1,
			Method:  "demo",
		}
		if err := enc.Encode(first); err != nil {
			_ = pw.CloseWithError(err)
			return
		}

		inputs := []string{"client input A", "client input B", "client input C"}
		for _, token := range inputs {
			if err := ctx.Err(); err != nil {
				_ = pw.CloseWithError(err)
				return
			}

			notif := clientInputNotification{
				JSONRPC: "2.0",
				Method:  "client/input",
				Params: map[string]interface{}{
					"token": token,
				},
			}
			if err := enc.Encode(notif); err != nil {
				_ = pw.CloseWithError(err)
				return
			}

			if err := sleepWithContext(ctx, 150*time.Millisecond); err != nil {
				_ = pw.CloseWithError(err)
				return
			}
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.URL, pr)
	if err != nil {
		log.Fatalf("create request failed err=%v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("request failed err=%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected http status=%d", resp.StatusCode)
	}

	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var tokenBuffer strings.Builder

	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			continue
		}

		var msg streamMessage
		if err := json.Unmarshal(line, &msg); err != nil {
			log.Fatalf("decode stream message failed err=%v raw=%s", err, string(line))
		}

		if len(msg.Error) > 0 {
			log.Fatalf("server returned error=%s", string(msg.Error))
		}

		if msg.Method == "notifications/token" && len(msg.Params) > 0 {
			var params struct {
				Token string `json:"token"`
			}
			if err := json.Unmarshal(msg.Params, &params); err != nil {
				log.Fatalf("decode params failed err=%v raw=%s", err, string(msg.Params))
			}
			tokenBuffer.WriteString(params.Token)
			fmt.Print(params.Token)
			continue
		}

		if msg.Method == "notifications/client_token" && len(msg.Params) > 0 {
			var params struct {
				Token string `json:"token"`
			}
			if err := json.Unmarshal(msg.Params, &params); err != nil {
				log.Fatalf("decode params failed err=%v raw=%s", err, string(msg.Params))
			}
			log.Printf("server ack client token=%s", params.Token)
			continue
		}

		if len(msg.ID) > 0 && len(msg.Result) > 0 {
			var result struct {
				Content       string `json:"content"`
				ClientContent string `json:"client_content"`
			}
			if err := json.Unmarshal(msg.Result, &result); err != nil {
				log.Fatalf("decode result failed err=%v raw=%s", err, string(msg.Result))
			}

			fmt.Print("\n")
			log.Printf("stream completed id=%s content=%s client_content=%s", string(msg.ID), result.Content, result.ClientContent)
			return
		}
	}

	if err := scanner.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("read stream failed err=%v", err)
	}
	log.Printf("no stream data received")
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	timer := time.NewTimer(d)
	defer func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
