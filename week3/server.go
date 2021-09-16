package week3

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// DefaultServerCloseSIG close信号量
var DefaultServerCloseSIG = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV}

type Service interface {
	Start() error
	ShutDown() error
}

type Server struct {
	Services []Service
}

// NewServer 实例化一个 server
func NewServer() *Server {
	return &Server{Services: make([]Service, 0, 2)}
}

// Register 注册一个 service
func (s *Server) Register(service Service) {
	s.Services = append(s.Services, service)
}

func (s *Server) Start() error {
	if len(s.Services) == 0 {
		fmt.Println("No service!")
		return nil
	}

	ch := make(chan os.Signal)

	for _, service := range s.Services {
		go func(srv Service) {
			e := srv.Start()
			if e != nil {
				fmt.Println("Error")
				ch <- syscall.SIGTERM
			}
		}(service)
	}

	signal.Notify(ch, DefaultServerCloseSIG...)

	sig := <-ch
	fmt.Printf("\nreceive sig: %s\n", sig.String())
	err := s.ShutDown()

	return err
}

func (s *Server) ShutDown() error {
	ctx, cf := context.WithTimeout(context.Background(), time.Second*5)
	defer cf()
	g, ctx := errgroup.WithContext(ctx)
	for _, service := range s.Services {
		g.Go(service.ShutDown)
	}

	ch := make(chan struct{}, 1)

	go func(ctx context.Context, ch chan struct{}) {
		select {
		case <-ch:
			fmt.Println("All services closed")
		case <-ctx.Done():
			fmt.Println("time out")
			os.Exit(-1)
		}
	}(ctx, ch)

	_ = g.Wait()
	ch <- struct{}{}

	return nil
}
