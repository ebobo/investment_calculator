package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebobo/investment_calculator_record/pkg/service"
	"github.com/jessevdk/go-flags"
)

var opt struct {
	GRPCServerAddr string `short:"g" long:"grpc-addr" default:":9092" description:"gRPC server address"`
}

func main() {
	_, err := flags.ParseArgs(&opt, os.Args)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	recordService := server.New(server.Config{
		GRPCListenAddr: opt.GRPCAddr,
	})

	e := service.Start()
	if e != nil {
		log.Fatalf("error starting server: %v", e)
	}

	// Block forever
	// Capture Ctrl-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	server.Shutdown()

}
