package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/defrox/enterprise-mcp-server/internal/audit"
	"github.com/defrox/enterprise-mcp-server/internal/auth"
	"github.com/defrox/enterprise-mcp-server/internal/tools"
)

type MCPRequest struct {
	Tool   string                 `json:"tool"`
	Role   string                 `json:"role"`
	Params map[string]interface{} `json:"params"`
}

func main() {
	siem := audit.NewSIEMLogger()

	http.HandleFunc("/mcp/execute", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req MCPRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json payload", http.StatusBadRequest)
			return
		}

		if !auth.IsAuthorized(req.Role, req.Tool) {
			siem.LogDecision(req.Role, req.Tool, false, "rbac_policy_violation")
			http.Error(w, "forbidden: role not authorized for tool", http.StatusForbidden)
			return
		}

		result, err := tools.ExecuteReadOnlyTool(req.Tool, req.Params)
		if err != nil {
			siem.LogDecision(req.Role, req.Tool, false, err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		siem.LogDecision(req.Role, req.Tool, true, "authorized_execution")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(result)
	})

	log.Println("Enterprise MCP Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
