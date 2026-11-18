package tools

import "testing"

func TestExecuteReadOnlyTool(t *testing.T) {
	t.Run("Valid query with explicit limit", func(t *testing.T) {
		res, err := ExecuteReadOnlyTool("query_analytics", map[string]interface{}{"limit": float64(25)})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if res.Status != "ok" || len(res.Records) != 1 {
			t.Fatalf("unexpected response structure: %+v", res)
		}
		if res.Records[0]["applied_cap"] != 25 {
			t.Errorf("expected applied_cap=25, got %v", res.Records[0]["applied_cap"])
		}
	})

	t.Run("Rejects limit above maximum cap", func(t *testing.T) {
		_, err := ExecuteReadOnlyTool("query_analytics", map[string]interface{}{"limit": float64(500)})
		if err == nil {
			t.Fatal("expected error for limit > 100, got nil")
		}
	})

	t.Run("Rejects unknown tool", func(t *testing.T) {
		_, err := ExecuteReadOnlyTool("delete_users", nil)
		if err == nil {
			t.Fatal("expected error for unknown tool, got nil")
		}
	})
}
