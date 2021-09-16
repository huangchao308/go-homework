package week3

import (
	"fmt"
	"time"
)

type CgiService struct {
	Name string
}

func NewCgiService() *CgiService {
	return &CgiService{Name: "cgi_service"}
}

func (cs *CgiService) Start() error {
	fmt.Println("starting cgi_service")
	time.Sleep(1 * time.Second)
	fmt.Println("cgi_service started")

	return nil
}

func (cs *CgiService) ShutDown() error {
	fmt.Println("begin to shut down cgi_service")
	time.Sleep(6 * time.Second)
	fmt.Println("cgi_service closed")

	return nil
}
