package internal

import "encoding/json"

type TestRequest map[string]interface{}

func NewTestRequest() TestRequest {
	return make(map[string]interface{})
}

func (r TestRequest) SetJsonrpc(v any) TestRequest {
	r["jsonrpc"] = v
	return r
}

func (r TestRequest) MockMethod() TestRequest {
	r["method"] = "mocked_method"
	return r
}

func (r TestRequest) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		panic("unexpected error on TestRequest.String(): " + err.Error())
	}
	return string(b)
}
