package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"

	"github.com/NethermindEth/juno/pkg/jsonrpc"

	"github.com/NethermindEth/juno/internal/rpc"
)

//go:generate mockgen -destination mock_blockmanager_test.go -source ./../../../db/block/manager.go -package internal BlockManager
//go:generate mockgen -destination mock_transactionmanager_test.go -source ./../../../db/transaction/manager.go -package internal TransactionManager

// TestRpc_InvalidHTTPMethods tests that the server returns a 405 error when an invalid HTTP method
// is used.
func TestRpc_InvalidHTTPMethods(t *testing.T) {
	var (
		rpcHandler = new(jsonrpc.JsonRpc)
		addr       = "localhost:6060"
		pattern    = "/rpc"
		url        = "http://" + addr + pattern
		rpcServer  = rpc.NewHttpRpc(addr, pattern, rpcHandler, nil)
	)

	invalidMethods := []string{
		http.MethodGet,
		http.MethodConnect,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPut,
		http.MethodTrace,
	}
	for _, invalidMethod := range invalidMethods {
		t.Run(invalidMethod, func(t *testing.T) {
			request, _ := http.NewRequest(invalidMethod, url, bytes.NewBufferString(`{"jsonrpc": "2.0", "method": "some_method", "id": 1}`))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			rpcServer.ServeHTTP(response, request)

			allowHeader := response.Header().Get("Allow")

			assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
			assert.Equal(t, http.MethodPost, allowHeader)
		})
	}
}

// TestRpc_ContentType tests that the server returns a 415 error when an invalid Content-type header
// is not present or different from application/json.
func TestRpc_ContentType(t *testing.T) {
	var (
		rpcHandler = new(jsonrpc.JsonRpc)
		addr       = "localhost:6060"
		pattern    = "/rpc"
		url        = "http://" + addr + pattern
		rpcServer  = rpc.NewHttpRpc(addr, pattern, rpcHandler, nil)
		requestStr = `{"jsonrpc": "2.0", "method": "some_method", "id": 1}`
	)

	tests := []struct {
		name               string
		contentType        string
		expectedStatusCode int
	}{
		{
			name:        "empty",
			contentType: "",
		},
		{
			name:        "invalid",
			contentType: "text/html",
		},
		{
			name:        "application/json with charset",
			contentType: "application/json; charset=utf-8",
		},
		{
			name:        "application/json with boundary",
			contentType: "application/json; boundary=something",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(requestStr)))
			request.Header.Set("Content-Type", tt.contentType)
			response := httptest.NewRecorder()
			rpcServer.ServeHTTP(response, request)

			assert.Equal(t, http.StatusUnsupportedMediaType, response.Code)
		})
	}
}

// TestRpc_InvalidRequestField_Jsonrpc tests that the server returns ErrInvalidRequest when the "jsonrpc"
// field is different from "2.0" on the JSON-RPC request.
func TestRpc_InvalidRequestField_Jsonrpc(t *testing.T) {
	var (
		rpcHandler = new(jsonrpc.JsonRpc)
		addr       = "localhost:6060"
		pattern    = "/rpc"
		url        = "http://" + addr + pattern
		rpcServer  = rpc.NewHttpRpc(addr, pattern, rpcHandler, nil)
	)

	tests := []struct {
		name        string
		requestBody string
		want        string
	}{
		{
			name:        "without jsonrpc field",
			requestBody: `{"method": "some_method", "id": 1}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "without jsonrpc value",
			requestBody: `{"jsonrpc": , "method": "some_method", "id": 1}`,
			want:        ErrParseError,
		},
		{
			name:        "empty string",
			requestBody: `{"jsonrpc": "", "method": "some_method", "id": 1}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "number",
			requestBody: `{"jsonrpc": 2.0, "method": "some_method", "id": 1}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "object",
			requestBody: `{"jsonrpc": {}, "method": "some_method", "id": 1}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "boolean",
			requestBody: `{"jsonrpc": true, "method": "some_method", "id": 1}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "null",
			requestBody: `{"jsonrpc": null, "method": "some_method", "id": 1}`,
			want:        ErrInvalidRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newRequest(url, tt.requestBody)
			response := httptest.NewRecorder()
			rpcServer.ServeHTTP(response, request)

			body, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, tt.want, string(body))
		})
	}
}

// TestRpc_InvalidRequestField_Id tests that the server returns ErrInvalidRequest when the type of
// the "id" field is different from "string" or "number".
func TestRpc_InvalidRequestField_Id(t *testing.T) {
	var (
		rpcHandler = new(jsonrpc.JsonRpc)
		addr       = "localhost:6060"
		pattern    = "/rpc"
		url        = "http://" + addr + pattern
		rpcServer  = rpc.NewHttpRpc(addr, pattern, rpcHandler, nil)
	)

	tests := []struct {
		name        string
		requestBody string
		want        string
	}{
		{
			name:        "without value",
			requestBody: `{"jsonrpc": "2.0", "method": "some_method", "id": }`,
			want:        ErrParseError,
		},
		{
			name:        "empty string",
			requestBody: `{"jsonrpc": "2.0", "method": "some_method", "id": ""}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "object",
			requestBody: `{"jsonrpc": "2.0", "method": "some_method", "id": {} }`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "boolean",
			requestBody: `{"jsonrpc": "2.0", "method": "some_method", "id": true }`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "float",
			requestBody: `{"jsonrpc": "2.0", "method": "some_method", "id": 1.0 }`,
			want:        ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newRequest(url, tt.requestBody)
			response := httptest.NewRecorder()
			rpcServer.ServeHTTP(response, request)

			body, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, tt.want, string(body))
		})
	}
}

// TestRpc_InvalidRequestField_Method tests that the server returns ErrInvalidRequest when the type of
// the "method" field is different from "string".
func TestRpc_InvalidRequestField_Method(t *testing.T) {
	var (
		rpcHandler = new(jsonrpc.JsonRpc)
		addr       = "localhost:6060"
		pattern    = "/rpc"
		url        = "http://" + addr + pattern
		rpcServer  = rpc.NewHttpRpc(addr, pattern, rpcHandler, nil)
	)

	tests := []struct {
		name        string
		requestBody string
		want        string
	}{
		{
			name:        "without value",
			requestBody: `{"jsonrpc": "2.0", "method": , "id": 1}`,
			want:        ErrParseError,
		},
		{
			name:        "object",
			requestBody: `{"jsonrpc": "2.0", "method": {}, "id": 1}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "boolean",
			requestBody: `{"jsonrpc": "2.0", "method": true, "id": 1}`,
			want:        ErrInvalidRequest,
		},
		{
			name:        "number",
			requestBody: `{"jsonrpc": "2.0", "method": 123, "id": 1}`,
			want:        ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newRequest(url, tt.requestBody)
			response := httptest.NewRecorder()
			rpcServer.ServeHTTP(response, request)

			body, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, tt.want, string(body))
		})
	}
}
