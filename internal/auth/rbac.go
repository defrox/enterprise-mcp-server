package auth

// rolePermissions maps each enterprise role to its allowed MCP tools.
var rolePermissions = map[string]map[string]bool{
	"admin": {
		"query_analytics": true,
		"system_health":   true,
	},
	"ai-analyst": {
		"query_analytics": true,
	},
}

// IsAuthorized checks whether a given role is permitted to execute the requested tool.
func IsAuthorized(role string, tool string) bool {
	allowedTools, exists := rolePermissions[role]
	if !exists {
		return false
	}
	return allowedTools[tool]
}
