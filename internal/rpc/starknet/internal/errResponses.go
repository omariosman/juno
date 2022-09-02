package internal

const (
	ErrInvalidRequest = `{"jsonrpc":"2.0","error":{"code":-32600,"message":"Invalid Request"}}`
	ErrParseError     = `{"jsonrpc":"2.0","error":{"code":-32700,"message":"Parse error"}}`
	ErrInvalidParams  = `{"jsonrpc":"2.0","error":{"code":-32602,"message":"Invalid params"},"id":1}`
	ErrInvalidBlockId = `{"jsonrpc":"2.0","error":{"Code":24,"Message":"Invalid block id"},"id":1}`
)
