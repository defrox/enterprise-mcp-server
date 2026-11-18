package tools

import (
	"errors"
	"fmt"
)

// QueryResult represents a structured read-only response returned to the AI agent.
type QueryResult struct {
	Tool    string                   `json:"tool"`
	Status  string                   `json:"status"`
	Records []map[string]interface{} `json:"records"`
}

// ExecuteReadOnlyTool dispatches validated, read-only enterprise tool requests.
func ExecuteReadOnlyTool(tool string, params map[string]interface{}) (QueryResult, error) {
	switch tool {
	case "query_analytics":
		limit := 10
		if rawLimit, ok := params["limit"].(float64); ok {
			if rawLimit <= 0 || rawLimit > 100 {
				return QueryResult{}, errors.New("parameter 'limit' must be between 1 and 100")
			}
			limit = int(rawLimit)
		}

		return QueryResult{
			Tool:   tool,
			Status: "ok",
			Records: []map[string]interface{}{
				{
					"metric":      "monthly_active_agents",
					"value":       142,
					"applied_cap": limit,
				},
			},
		}, nil

	case "system_health":
		return QueryResult{
			Tool:   tool,
			Status: "healthy",
			Records: []map[string]interface{}{
				{"gateway": "operational", "mode": "read-only"},
			},
		}, nil

	default:
		return QueryResult{}, fmt.Errorf("unknown or unsupported tool: %s", tool)
	}
}
