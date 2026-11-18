package auth

import "testing"

func TestIsAuthorized(t *testing.T) {
	tests := []struct {
		name     string
		role     string
		tool     string
		expected bool
	}{
		{"Admin can query analytics", "admin", "query_analytics", true},
		{"Admin can check system health", "admin", "system_health", true},
		{"AI Analyst can query analytics", "ai-analyst", "query_analytics", true},
		{"AI Analyst cannot check system health", "ai-analyst", "system_health", false},
		{"Unknown role is denied by default", "guest", "query_analytics", false},
		{"Empty role is denied", "", "query_analytics", false},
		{"Unknown tool is denied", "admin", "drop_database", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsAuthorized(tc.role, tc.tool)
			if got != tc.expected {
				t.Errorf("IsAuthorized(%q, %q) = %v; want %v", tc.role, tc.tool, got, tc.expected)
			}
		})
	}
}
