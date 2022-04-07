package goroutine

import (
	"context"
	"fmt"
	"net/http"
)

func Start(server *http.Server, path string) error {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my web server!")
	})
	return server.ListenAndServe()
}

func Shutdown(ctx context.Context, server *http.Server) error {
	return server.Shutdown(ctx)
}
