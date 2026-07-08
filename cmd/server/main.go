package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type MCPRequest struct {
	Tool   string                 `json:"tool"`
	Role   string                 `json:"role"`
	Params map[string]interface{} `json:"params"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	http.HandleFunc("/mcp/execute", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req MCPRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		if req.Role != "ai-analyst" && req.Role != "admin" {
			logger.Warn("unauthorized_mcp_access", "tool", req.Tool, "role", req.Role)
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		logger.Info("mcp_tool_invoked",
			"tool", req.Tool,
			"role", req.Role,
			"timestamp", time.Now().UTC().Format(time.RFC3339),
		)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "executed",
			"tool":   req.Tool,
		})
	})

	logger.Info("Enterprise MCP Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Error("server_failed", "error", err)
		os.Exit(1)
	}
}
