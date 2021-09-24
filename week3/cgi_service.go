package week3

import (
	"errors"
	"fmt"
	"time"
)

type CgiService struct {
	Name string
}

func NewCgiService() CgiService {
	return CgiService{Name: "cgi_service"}
}

func (cs *CgiService) Start() error {
	fmt.Println("starting cgi_service")
	time.Sleep(1 * time.Second)

	return errors.New("test")
}

func (cs *CgiService) Close(ch chan struct{}) error {
	fmt.Println("begin to shut down cgi_service")
	err := cs.DoBeforeClose()
	if err != nil {
		return err
	}
	ch <- struct{}{}

	return nil
}

func (cs *CgiService) DoBeforeClose() error {
	time.Sleep(1 * time.Second)
	return errors.New("test")
}
