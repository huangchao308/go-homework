package week3

import (
	"fmt"
	"time"
)

type LogService struct {
	Name string
}

func NewLogService() *LogService {
	return &LogService{Name: "log_service"}
}

func (cs *LogService) Start() error {
	fmt.Println("starting log_service")
	time.Sleep(1 * time.Second)
	fmt.Println("log_service started")

	return nil
}

func (cs *LogService) ShutDown() error {
	fmt.Println("begin to shut down log_service")
	time.Sleep(1 * time.Second)
	fmt.Println("log_service closed")

	return nil
}
