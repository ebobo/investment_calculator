package server

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// numConnectRetries is the number of times we try to connect a gRPC client
	// to the server when ensuring the server has come up before we give up.
	numConnectRetries = 10

	// delayBetweenConnectRetries is how long we wait between each connection
	// attempt when verifying that the server is up.
	delayBetweenConnectRetries = (100 * time.Millisecond)
)

// startGRPC takes care of starting the gRPC server.  Note that this has to happen
// before you call startHTTP if you are using RESTMuxViaGRPC since it depends on
// the gRPC server being up before it can start safely.
func (s *Server) startGRPC() error {
	listener, err := net.Listen("tcp", s.grpcListenAddr)
	if err != nil {
		return err
	}

	grpcServer := s.loanService.GRPCserver()

	// Set up graceful shutdown
	go func() {
		<-s.ctx.Done()
		grpcServer.GracefulStop()
	}()

	// Start gRPC server
	go func() {
		log.Printf("starting gRPC interface '%s'", s.grpcListenAddr)
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Printf("gRPC interface '%s' down, %v", s.grpcListenAddr, err)
		} else {
			log.Printf("gRPC interface '%s' down", s.grpcListenAddr)
		}
		s.grpcStopped.Done()
	}()

	// Ensure the gRPC server is really up before we return
	connectSuccess := false
	for i := 0; i < numConnectRetries; i++ {
		client, err := grpc.Dial(s.grpcListenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err == nil {
			connectSuccess = true
			client.Close()
			break
		}

		log.Printf("waiting for gRPC endpoint '%s' to come up (attempt %d)", s.grpcListenAddr, i+1)
		time.Sleep(delayBetweenConnectRetries)
	}

	if !connectSuccess {
		return fmt.Errorf("failed to connect to gRPC endpoint %s to check for liveness", s.grpcListenAddr)
	}

	s.grpcStarted.Done()
	return nil
}
