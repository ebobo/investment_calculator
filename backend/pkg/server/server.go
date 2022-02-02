package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/ebobo/investment_calculator/pkg/service"
)

// Server takes care of instantiating and running service and other dependencies.
type Server struct {
	grpcListenAddr string
	httpListenAddr string
	msGRPCAddr     string
	loanService    *service.LoanServer
	grpcStarted    *sync.WaitGroup
	grpcStopped    *sync.WaitGroup
	httpStarted    *sync.WaitGroup
	httpStopped    *sync.WaitGroup
	ctx            context.Context
	cancel         context.CancelFunc
}

// Config is the server configuration
type Config struct {
	GRPCListenAddr string
	HTTPListenAddr string
	MSGPRCAddr     string
}

func New(c Config) *Server {
	return &Server{
		grpcListenAddr: c.GRPCListenAddr,
		httpListenAddr: c.HTTPListenAddr,
		msGRPCAddr:     c.MSGPRCAddr,
		grpcStarted:    &sync.WaitGroup{},
		grpcStopped:    &sync.WaitGroup{},
		httpStarted:    &sync.WaitGroup{},
		httpStopped:    &sync.WaitGroup{},
	}
}

func (s *Server) Start() error {
	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.loanService = service.NewLoanServer(s.ctx)

	// Start gRPC interface.
	s.grpcStarted.Add(2)
	s.grpcStopped.Add(1)
	err := s.startGRPC()
	if err != nil {
		return err
	}

	go func() {
		err := s.connectMSGRPC()
		if err != nil {
			fmt.Println(err)
		}
		s.grpcStarted.Done()
	}()
	s.grpcStarted.Wait()

	// Start the HTTP interface
	s.httpStarted.Add(1)
	s.httpStopped.Add(1)
	err = s.startHTTP()
	if err != nil {
		return err
	}
	s.httpStarted.Wait()

	return nil
}

func (s *Server) Shutdown() {
	if s.cancel != nil {
		s.cancel()
	}
	s.httpStopped.Wait()
	s.grpcStopped.Wait()
}
