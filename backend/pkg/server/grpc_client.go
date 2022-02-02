package server

import (
	"log"

	"github.com/ebobo/investment_calculator/pkg/api/go/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// the gRPC server being up before it can start safely.
func (s *Server) connectMSGRPC() error {

	conn, err := grpc.Dial(s.msGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	} else {
		log.Printf("connnect to gRPC server at %s", s.msGRPCAddr)
	}

	s.loanService.MsClient = proto.NewRecordServiceClient(conn)
	return nil
}
