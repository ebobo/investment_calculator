package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ebobo/investment_calculator/pkg/api/go/proto/v1"
	"github.com/ebobo/investment_calculator/pkg/utility"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

// LoanService that implements the pb API
type LoanServer struct {
	recordStream   proto.InvestmentService_SaveRecordServer
	ctx            context.Context
	grpcClientConn *grpc.ClientConn
	proto.UnimplementedInvestmentServiceServer
}

// NewUserServer creates a new Service instance
func NewLoanServer(c context.Context) *LoanServer {
	return &LoanServer{}
}

// Run grpc server
func (s *LoanServer) GRPCserver() *grpc.Server {
	gs := grpc.NewServer()
	//gRPC Server Reflection provides information about publicly-accessible gRPC services on a server
	reflection.Register(gs)
	proto.RegisterInvestmentServiceServer(gs, s)

	return gs
}

// CloseInternalClientConn closes the internal client connection if we use the RESTMuxViaGRPC()
// style calls.  While this isn't strictly necessary, being able to shut it down will clean things
// up so unit tests don't complain about leaked goroutines.
func (server *LoanServer) CloseInternalClientConn() error {
	if server.grpcClientConn != nil {
		return server.grpcClientConn.Close()
	}
	return nil
}

//GetResult implementation
func (s *LoanServer) GetResult(ctx context.Context, in *proto.Case) (*proto.Report, error) {

	// fmt.Println(in.GetHouseValue(), in.GetInterestRate(), in.GetPaymentYear(), in.GetPeriodicFee(), in.GetOneTimeFee())

	periodocPaymentAmount, totalPayment, totalIntersetAndFees := utility.BankLoanEqualInstallments(
		uint32(in.GetHouseValue()),
		in.GetInterestRate(),
		uint32(in.GetPaymentYear()),
		float32(in.GetPeriodicFee()),
		float32(in.GetOneTimeFee()))

	report := &proto.Report{Client: in.GetClient(), TotalInterest: totalIntersetAndFees, PeriodicPayment: periodocPaymentAmount, TotalPayment: totalPayment}
	err := s.recordStream.Send(report)
	if err != nil {
		fmt.Errorf("Error %v \n", err)
	}
	return report, nil
}

//SaveRecord implementation
func (s *LoanServer) SaveRecord(_ *emptypb.Empty, stream proto.InvestmentService_SaveRecordServer) error {
	s.recordStream = stream
	<-s.ctx.Done()
	return nil
}

//GetRecords implementation
// func (s *LoanServer) GetRecords(ctx context.Context, in *proto.User) (*proto.Records, error) {
// report := &proto.Reports{proto.Report{}}
// return report, nil
// }

// RESTMuxViaGRPC creates a mux which uses an internal gRPC client.  This has the benefit that
// REST accesses will go through interceptors.
func (s *LoanServer) RESTMuxViaGRPC(ctx context.Context, listenAddr string) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	// Get the port
	parts := strings.Split(listenAddr, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("cannot get port from listenAddr '%s'", listenAddr)
	}
	port, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing port number in listenAddr %s: %w", listenAddr, err)
	}

	// Create the client
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("unable to wire up internal REST client to gRPC interface: %w", err)
	}

	s.grpcClientConn = conn

	// // Register properties handler
	if err := proto.RegisterInvestmentServiceHandler(ctx, mux, conn); err != nil {
		return nil, fmt.Errorf("RegisterPropertiesHandler failed: %w", err)
	}
	return mux, nil
}

// RESTMux that uses the service directly.  This means requests will not go through interceptors.
func (s *LoanServer) RESTMux(ctx context.Context) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	if err := proto.RegisterInvestmentServiceHandlerServer(ctx, mux, s); err != nil {
		return nil, fmt.Errorf("RegisterUsersHandler failed: %w", err)
	}

	return mux, nil
}
