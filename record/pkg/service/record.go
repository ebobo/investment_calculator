package service

import (
	"context"
	"io"
	"log"

	"github.com/ebobo/investment_calulator_record/pkg/api/go/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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
	} else {
		log.Printf("connnect to grpc server at %s", ms.grpcServerAddr)
	}

	c := proto.NewInvestmentServiceClient(conn)

	stream, err := c.SaveRecord(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("EOF")
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Println("Record : ", resp.Client, resp.TotalInterest, resp.PeriodicPayment, resp.TotalPayment)
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}
