package main

import (
	"context"
	"log"

	worker "learn-go/goroutine"
	"learn-go/handle_error"
)

func main() {
	ctx := context.Background()

	// Week 2 - Error Handling
	service := handle_error.NewStudentService()
	_, err := service.GetByID(1)
	if err != nil {
		log.Printf("stack trace: \n%+v\n", err)
	}

	//	Week 3 - Goroutine
	worker.Run(ctx)
}
