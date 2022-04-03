package goroutine

import (
	"context"
	"io"
	"net/http"
)

func HTTPServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Successful")
}

func Start(server *http.Server, path string) error {
	http.HandleFunc(path, HTTPServer)
	return server.ListenAndServe()
}

func Shutdown(ctx context.Context, server *http.Server) error {
	return server.Shutdown(ctx)
}
