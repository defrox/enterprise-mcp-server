# Enterprise MCP Server Blueprint (Go)

[![Enterprise MCP Server CI](https://github.com/defrox/enterprise-mcp-server/actions/workflows/ci.yml/badge.svg)](https://github.com/defrox/enterprise-mcp-server/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Security](https://img.shields.io/badge/Security-RBAC%20%2B%20SIEM-2E7D32?logo=owasp&logoColor=white)](https://owasp.org/)

A production-grade **Model Context Protocol (MCP)** server written in **Go**, engineered for enterprise environments requiring strict **Role-Based Access Control (RBAC)**, immutable structured audit logging, and low-latency tool execution for autonomous AI agents.

## Architecture

```mermaid
sequenceDiagram
    participant Agent as AI Agent (LangGraph / Claude)
    participant MCP as Enterprise MCP Server (Go)
    participant SIEM as SIEM / Audit Log
    participant DB as Enterprise Data Source

    Agent->>MCP: POST /mcp/execute (Tool + Role + Params)
    MCP->>MCP: Validate RBAC Policy
    alt Unauthorized Role
        MCP->>SIEM: Log Security Alert (JSON)
        MCP-->>Agent: 403 Forbidden
    else Authorized Role
        MCP->>SIEM: Log Immutable Audit Trail
        MCP->>DB: Execute Read-Only Query
        DB-->>MCP: Result Payload
        MCP-->>Agent: 200 OK (Structured Response)
    end
```

## Key Features
* **Zero-Trust RBAC Enforcement:** Blocks unauthorized tool invocations at the gateway boundary before touching downstream data systems.
* **SIEM-Ready Structured Logging:** Outputs JSON audit events via Go's `log/slog` with tool name, caller role, and UTC timestamps.
* **Cloud-Native Footprint:** Minimal memory overhead and multi-stage distroless Docker build ready for Kubernetes.

## Quickstart

```bash
# Run locally
go run ./cmd/server

# Test an authorized tool execution
curl -X POST http://localhost:8080/mcp/execute \
  -H "Content-Type: application/json" \
  -d '{"tool": "query_analytics", "role": "ai-analyst", "params": {"limit": 10}}'
```

## Author
**Christian Axel Schmidt** — Principal AI & Systems Architect
