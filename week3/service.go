package week3

import "context"

type DefaultService struct {
	Name   string
	ctx    context.Context
	cancel context.CancelFunc
}

func NewDefaultService(name string) *DefaultService {
	ctx, cancel := context.WithCancel(context.Background())
	return &DefaultService{
		Name:   name,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *DefaultService) Start() error {
	return nil
}

func (s *DefaultService) Close(ch chan struct{}) error {
	ch <- struct{}{}
	return nil
}
