package main

import (
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

// JSONRPCRequest 是一个最小化的 JSON-RPC 2.0 请求结构，用于演示 StreamableHTTP。
type JSONRPCRequest struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      interface{}            `json:"id"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

// JSONRPCNotification 是一个最小化的 JSON-RPC 2.0 通知（无 ID），用于流式推送中间结果。
type JSONRPCNotification struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

// JSONRPCResponse 是一个最小化的 JSON-RPC 2.0 响应结构。
type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonRPCNotification struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type clientInputParams struct {
	Token string `json:"token"`
}

type serverConfig struct {
	Addr       string
	TokenDelay time.Duration
}

func newServerConfig() serverConfig {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port != "" && strings.HasPrefix(port, ":") {
		return serverConfig{Addr: port, TokenDelay: 200 * time.Millisecond}
	}
	if port != "" {
		return serverConfig{Addr: ":" + port, TokenDelay: 200 * time.Millisecond}
	}

	addr := strings.TrimSpace(os.Getenv("MCP_ADDR"))
	if addr == "" {
		addr = ":9090"
	}
	return serverConfig{Addr: addr, TokenDelay: 200 * time.Millisecond}
}

func isJSONContentType(contentType string) bool {
	contentType = strings.TrimSpace(contentType)
	return strings.HasPrefix(strings.ToLower(contentType), "application/json")
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

func mcpHandler(cfg serverConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 允许 CORS 预检，方便用浏览器/脚本快速试验
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !isJSONContentType(r.Header.Get("Content-Type")) {
			http.Error(w, "content-type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		dec := json.NewDecoder(r.Body)

		var req JSONRPCRequest
		if err := dec.Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("invalid json: %v", err), http.StatusBadRequest)
			return
		}
		log.Printf("received request id=%v method=%s", req.ID, req.Method)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Accel-Buffering", "no")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}

		enc := json.NewEncoder(w) // 每次 Encode 都会写入换行，适合 JSON Lines
		ctx := r.Context()

		clientTokens := make(chan string, 16)
		readErr := make(chan error, 1)
		go func() {
			defer close(clientTokens)
			for {
				var notif jsonRPCNotification
				if err := dec.Decode(&notif); err != nil {
					if errors.Is(err, io.EOF) {
						readErr <- nil
						return
					}
					readErr <- err
					return
				}

				if notif.Method != "client/input" {
					continue
				}

				var params clientInputParams
				if err := json.Unmarshal(notif.Params, &params); err != nil {
					readErr <- err
					return
				}

				select {
				case clientTokens <- params.Token:
				case <-ctx.Done():
					readErr <- ctx.Err()
					return
				}
			}
		}()

		serverTokens := []string{"Hello", ", ", "StreamableHTTP", " ", "demo", "!"}
		var clientContent strings.Builder

		ticker := time.NewTicker(cfg.TokenDelay)
		defer ticker.Stop()

		serverTokenIndex := 0
		readDone := false
		for {
			if err := ctx.Err(); err != nil {
				log.Printf("client disconnected id=%v err=%v", req.ID, err)
				return
			}

			serverDone := serverTokenIndex >= len(serverTokens)
			if serverDone && readDone {
				break
			}

			select {
			case token, ok := <-clientTokens:
				if !ok {
					readDone = true
					continue
				}
				clientContent.WriteString(token)
				log.Printf("received client input id=%v token=%q", req.ID, token)

				fromClient := JSONRPCNotification{
					JSONRPC: "2.0",
					Method:  "notifications/client_token",
					Params: map[string]interface{}{
						"token": token,
					},
				}
				if err := enc.Encode(fromClient); err != nil {
					log.Printf("write notification failed id=%v err=%v", req.ID, err)
					return
				}
				flusher.Flush()

			case err := <-readErr:
				if err != nil {
					log.Printf("read request stream failed id=%v err=%v", req.ID, err)
					return
				}
				readDone = true

			case <-ticker.C:
				if serverTokenIndex >= len(serverTokens) {
					continue
				}
				token := serverTokens[serverTokenIndex]

				tokenNotif := JSONRPCNotification{
					JSONRPC: "2.0",
					Method:  "notifications/token",
					Params: map[string]interface{}{
						"token": token,
						"index": serverTokenIndex,
					},
				}
				if err := enc.Encode(tokenNotif); err != nil {
					log.Printf("write notification failed id=%v err=%v", req.ID, err)
					return
				}
				flusher.Flush()
				serverTokenIndex++
			}
		}

		finalResp := JSONRPCResponse{
			JSONRPC: "2.0",
			ID:      req.ID,
			Result: map[string]interface{}{
				"content":        strings.Join(serverTokens, ""),
				"client_content": clientContent.String(),
			},
		}
		if err := enc.Encode(finalResp); err != nil {
			log.Printf("write final response failed id=%v err=%v", req.ID, err)
			return
		}
		flusher.Flush()
		log.Printf("completed request id=%v", req.ID)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"mcp":    "StreamableHTTP demo",
	}); err != nil {
		log.Printf("write health response failed err=%v", err)
	}
}

func main() {
	cfg := newServerConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/mcp", mcpHandler(cfg))

	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Printf("starting server addr=%s", cfg.Addr)
	log.Printf("endpoint mcp=http://localhost%s/mcp", cfg.Addr)
	log.Printf("endpoint health=http://localhost%s/health", cfg.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed err=%v", err)
	}
}
