package goroutine

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context) {
	ctx, done := context.WithCancel(ctx)
	errGroup, groupCtx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

		select {
		case sig := <-signalChan:
			fmt.Printf("received signal: %s\n", sig)
			done()
		case <-groupCtx.Done():
			fmt.Println("done")
			return groupCtx.Err()
		}

		return nil
	})

	// HTTP Server
	server := &http.Server{Addr: ":8080"}

	errGroup.Go(func() error {
		fmt.Println("start http server")
		return Start(server, "/hello")
	})

	errGroup.Go(func() error {
		<-groupCtx.Done()
		fmt.Println("stop http server")
		return Shutdown(groupCtx, server)
	})

	// Force Stop
	time.AfterFunc(10*time.Second, func() {
		fmt.Println("force stop after 10s")
		done()
	})

	err := errGroup.Wait()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Println("context was cancelled")
		} else {
			fmt.Println("something went wrong")
		}
	} else {
		fmt.Println("completed")
	}
}
