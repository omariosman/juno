package rpc

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/NethermindEth/juno/pkg/jsonrpc"
)

type HttpRpc interface {
	http.Handler
	Run(chan<- error)
	Close(duration time.Duration) error
}

type httpRpc struct {
	server  *http.Server
	address string
	pattern string
	rpc2    *jsonrpc.JsonRpc
	logger  *zap.SugaredLogger
}

func NewHttpRpc(addr, pattern string, rpc *jsonrpc.JsonRpc, logger *zap.SugaredLogger) HttpRpc {
	return &httpRpc{
		address: addr,
		pattern: pattern,
		rpc2:    rpc,
		logger:  logger,
	}
}

func (h *httpRpc) Run(errCh chan<- error) {
	serverMux := http.NewServeMux()
	serverMux.Handle(h.pattern, h)
	h.server = &http.Server{Addr: h.address, Handler: serverMux}
	go func() {
		h.logger.Info("Listening for JSON-RPC connections on ", h.address)
		if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- errors.New("Failed to ListenAndServe on RPC server: " + err.Error())
		}
		close(errCh)
	}()
}

func (h *httpRpc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := h.rpc2.Call(body)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *httpRpc) Close(timeout time.Duration) error {
	h.logger.Info("Shutting down JSON-RPC server...")
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return h.server.Shutdown(ctx)
}
