package service

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/ebobo/investment_calulator_record/db"
	"github.com/ebobo/investment_calulator_record/pkg/api/go/proto/v1"
	"github.com/ebobo/investment_calulator_record/pkg/model"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	_ "github.com/mattn/go-sqlite3"
)

// RecordService to save record data to database
type RecordService struct {
	icGrpcServerAddr string
	msGrpcServerAddr string
	sqliteDatabase   *sqlx.DB
}

func New(icAddress string, msAddress string) *RecordService {
	return &RecordService{
		icGrpcServerAddr: icAddress,
		msGrpcServerAddr: msAddress,
		sqliteDatabase:   nil,
	}
}

//GetSavedRecords implementation
func (ms *RecordService) GetSavedRecords(ctx context.Context, in *proto.User) (*proto.Records, error) {

	reports := &proto.Records{}
	results, err := db.GetRecordsByClient(ms.sqliteDatabase, in.Client)

	if err != nil {
		return reports, err
	}
	for _, r := range results {
		reports.Reports = append(reports.Reports, &proto.Report{Client: r.Client, TotalInterest: r.TotalInterest, PeriodicPayment: r.PeriodicPayment, TotalPayment: r.TotalPayment})
	}
	log.Println("Records: ", reports)

	return reports, nil
}

// Run runs whole algorithm to process maps
func (ms *RecordService) Run() {
	log.Println("Running record service")

	err := db.CreateDataBase()
	if err != nil {
		log.Fatalf("can not create database %v", err)
	}

	// Open the created SQLite File
	sqlite, err := sqlx.Open("sqlite3", "../data/ic-database.db")
	if err != nil {
		log.Fatalf("can not open database %v", err)
	}

	ms.sqliteDatabase = sqlite

	// Defer Closing the database
	defer ms.sqliteDatabase.Close()

	e := db.CreateSchema(ms.sqliteDatabase) // Create Database Tables
	if e != nil {
		log.Fatalf("can not create schema  %v", e)
	}

	conn, err := grpc.Dial(ms.icGrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	} else {
		log.Printf("connect to grpc server at %s", ms.icGrpcServerAddr)
	}

	//Start a grpc server
	go func() {
		ms.startGRPC()
	}()

	c := proto.NewInvestmentServiceClient(conn)

	// WaitForReady configures the action to take when an RPC is attempted on broken connections or unreachable servers.
	// If waitForReady is false and the connection is in the TRANSIENT_FAILURE state, the RPC will fail immediately.
	// If waitForReady is true, the RPC client will block the call until a connection is available
	// (or the call is canceled or times out) and will retry the call if it fails due to a transient error.
	// gRPC will not retry if data was written to the wire unless the server indicates it did not process the data.
	stream, err := c.SaveRecord(context.Background(), &emptypb.Empty{}, grpc.WaitForReady(true))
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

			db.Addrecord(ms.sqliteDatabase, &model.Report{Client: resp.Client, TotalInterest: resp.TotalInterest, PeriodicPayment: resp.PeriodicPayment, TotalPayment: resp.TotalPayment})
		}
	}()

	<-done //we will wait until all response is received
	log.Printf("finished")
}

func (ms *RecordService) startGRPC() error {
	listener, err := net.Listen("tcp", ms.msGrpcServerAddr)

	if err != nil {
		return err
	}
	gs := grpc.NewServer()

	proto.RegisterRecordServiceServer(gs, ms)

	// Start gRPC server

	log.Printf("starting gRPC interface %s", ms.msGrpcServerAddr)
	e := gs.Serve(listener)

	return e

}
