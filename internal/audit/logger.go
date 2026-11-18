package audit

import (
	"log/slog"
	"os"
	"time"
)

// SIEMLogger wraps structured JSON logging for security and compliance auditing.
type SIEMLogger struct {
	logger *slog.Logger
}

// NewSIEMLogger initializes a JSON logger writing to stdout.
func NewSIEMLogger() *SIEMLogger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return &SIEMLogger{
		logger: slog.New(handler),
	}
}

// LogDecision records every tool execution attempt, outcome, and caller identity.
func (s *SIEMLogger) LogDecision(role string, tool string, allowed bool, reason string) {
	ts := time.Now().UTC().Format(time.RFC3339)
	if !allowed {
		s.logger.Warn("mcp_access_denied",
			"event_type", "security_alert",
			"role", role,
			"tool", tool,
			"reason", reason,
			"timestamp_utc", ts,
		)
		return
	}

	s.logger.Info("mcp_tool_executed",
		"event_type", "audit_trail",
		"role", role,
		"tool", tool,
		"reason", reason,
		"timestamp_utc", ts,
	)
}
