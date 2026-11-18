FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o mcp-server ./cmd/server

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/mcp-server /mcp-server
EXPOSE 8080
ENTRYPOINT ["/mcp-server"]
