// Package server provides MCP (Model Control Protocol) server implementations.
// copied from https://github.com/mark3labs/mcp-go/tree/v0.20.0/server/server.go
package server

import "github.com/mark3labs/mcp-go/mcp"

func createErrorResponse(
	id interface{},
	code int,
	message string,
) mcp.JSONRPCMessage {
	return mcp.JSONRPCError{
		JSONRPC: mcp.JSONRPC_VERSION,
		ID:      id,
		Error: struct {
			Code    int         `json:"code"`
			Message string      `json:"message"`
			Data    interface{} `json:"data,omitempty"`
		}{
			Code:    code,
			Message: message,
		},
	}
}
