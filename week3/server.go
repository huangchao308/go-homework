package week3

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// DefaultServerCloseSIG close信号量
var DefaultServerCloseSIG = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV}

type Service interface {
	Start() error
	Close(c chan error) error
}

type Server struct {
	Services map[string]Service
}

// NewServer 实例化一个 server
func NewServer() *Server {
	return &Server{Services: make(map[string]Service)}
}

// Register 注册一个 service
func (s *Server) Register(name string, service Service) {
	_, ok := s.Services[name]
	if !ok {
		s.Services[name] = service
	}
}

func (s *Server) Start() error {
	if len(s.Services) == 0 {
		fmt.Println("No service!")
		return nil
	}

	ch := make(chan os.Signal)
	g, _ := errgroup.WithContext(context.Background())
	for _, service := range s.Services {
		go func(srv Service) {
			g.Go(srv.Start)
		}(service)
	}
	signal.Notify(ch, DefaultServerCloseSIG...) //nolint:govet
	err := g.Wait()
	if err != nil {
		fmt.Printf("err: %+v", err)
		go func() {
			ch <- syscall.SIGTERM
		}()
	}

	sig := <-ch
	fmt.Printf("\nreceive sig: %s\n", sig.String())
	err = s.Close()
	if err != nil {
		fmt.Printf("Close services err: %+v", err)
	}

	return err
}

func (s *Server) Close() error {
	var err error
	ctx, cf := context.WithTimeout(context.Background(), time.Second*60)
	defer cf()
	wg := sync.WaitGroup{}
	for name, service := range s.Services {
		wg.Add(1)
		go func(n string, srv Service) {
			defer wg.Done()
			ch := make(chan error, 1)
			go func() {
				e := srv.Close(ch)
				if err == nil && e != nil {
					err = e
				}
			}()
			select {
			case e := <-ch:
				if e != nil {
					fmt.Printf("\nclose service %s err %v\n", n, e)
				} else {
					fmt.Printf("\nclose service %s succeed\n", n)
				}
			case <-ctx.Done():
				fmt.Printf("\nclose service %s %v.\n", n, ctx.Err())
			}
		}(name, service)
	}
	wg.Wait()
	fmt.Println("server closed!")
	return err
}
