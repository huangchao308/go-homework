package week3

import (
	"fmt"
	"time"
)

type LogService struct {
	Name string
}

func NewLogService() LogService {
	return LogService{Name: "log_service"}
}

func (ls *LogService) Start() error {
	fmt.Println("starting log_service")
	time.Sleep(1 * time.Second)
	fmt.Println("log_service started")

	return nil
}

func (ls *LogService) Close(ch chan error) error {
	fmt.Println("begin to shut down log_service")
	err := ls.DoBeforeClose()
	ch <- err

	return err
}

func (ls *LogService) DoBeforeClose() error {
	time.Sleep(10 * time.Second)
	return nil
}
