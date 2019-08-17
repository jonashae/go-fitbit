package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// CallbackServer starts a temporary HTTP server and awaits an OAuth2 callback
type CallbackServer struct {
	context context.Context
	server  *http.Server
	code    chan string
}

// NewCallbackServer creates a new CallbackServer
func NewCallbackServer(port int) *CallbackServer {
	handler := &CallbackServer{code: make(chan string)}

	address := fmt.Sprintf(":%d", port)
	server := &http.Server{Addr: address, Handler: handler}
	handler.server = server

	return handler
}

// AwaitCode waits for a token callback and returns it unless the context expires
func (callback *CallbackServer) AwaitCode(context context.Context) (string, error) {
	callback.context = context

	go callback.listen()
	defer callback.shutdown()

	select {
	case code := <-callback.code:
		return code, nil
	case <-context.Done():
		return "", errors.New("context timeout")
	}
}

func (callback *CallbackServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	callback.code <- r.URL.Query().Get("code")
}

func (callback *CallbackServer) listen() {
	err := callback.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func (callback *CallbackServer) shutdown() error {
	return callback.server.Shutdown(callback.context)
}
