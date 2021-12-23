package tools

import (
	"log"
	"os"
)

var logger *log.Logger

func GetLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stdout, "training-go-invoices ", log.LstdFlags)
	}
	return logger
}
