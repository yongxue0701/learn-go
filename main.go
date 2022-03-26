package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"learn-go/handle_error"
	"os"
)

var log = logrus.New()

func main() {
	log.Out = os.Stdout

	service := handle_error.NewStudentService()
	_, err := service.GetByID(1)
	if err != nil {
		log.Error(fmt.Printf("stack trace: \n%+v\n", err))
	}
}
