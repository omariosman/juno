package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"

	. "github.com/NethermindEth/juno/internal/log"
)

// ErrResponse represents the JSON formatted error response that is
// returned by the feeder gateway.
type ErrResponse struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Problems string `json:"problems,omitempty"`
}

// clientErr sets a 400 Bad Request header and then serves a JSON
// formatted client error. If msg != "", ErrResponse.Code will be set to
// starkErr ErrResponse.Message == msg, otherwise a generic error is
// returned.
func clientErr(w http.ResponseWriter, code int, starkErr, msg string) {
	var res *ErrResponse
	switch {
	case msg != "":
		res = &ErrResponse{Code: starkErr, Message: msg}
	default:
		res = &ErrResponse{
			Code:     strconv.Itoa(code),
			Message:  fmt.Sprintf("%d: %s", code, http.StatusText(code)),
			Problems: http.StatusText(code),
		}
	}

	// XXX: Might prefer json.Marshal instead to reduce the payload size.
	raw, err := json.MarshalIndent(&res, "", "  ")
	if err != nil {
		Logger.Errorw("failed to marshal JSON", "error", err.Error())
		serverErr(w, err)
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(raw)
}

// notImplementedErr sets a 501 Not Implemented header and then serves
// a JSON formatted error.
func notImplementedErr(w http.ResponseWriter) {
	clientErr(w, http.StatusNotImplemented, "", "")
}

// serverErr sets a 500 Internal Server Error header and then serves a
// JSON formatted client error.
func serverErr(w http.ResponseWriter, err error) {
	// XXX: The server will print a stack trace when a panic occurs but in
	// an unstructured manner. I suspect this is because the the zap
	// logger cannot be used as a http.Server.ErrorLog. An alternative
	// could involve instantiating the logger with a log.Logger that
	// outputs to io.Discard 🤔.
	Logger.Errorf("%s\n%s", err.Error(), debug.Stack())

	res := &ErrResponse{
		Code:     strconv.Itoa(http.StatusInternalServerError),
		Message:  http.StatusText(http.StatusInternalServerError),
		Problems: http.StatusText(http.StatusInternalServerError),
	}

	// XXX: Might prefer json.Marshal instead to reduce the payload size.
	raw, err := json.MarshalIndent(&res, "", "  ")
	if err != nil {
		Logger.Errorw("failed to marshal JSON", "error", err.Error())
		panic(err)
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(raw)
}
