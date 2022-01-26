package service

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RecordService to save record data to database
type RecordService struct {
	grpcServerAddr string
}

func New(address string) *RecordService {
	return &RecordService{
		grpcServerAddr: address,
	}
}

// Run runs whole algorithm to process maps
func (ms *RecordService) Run() {
	log.Println("Running record service")

	conn, err := grpc.Dial(ms.grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	c := proto.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

}
