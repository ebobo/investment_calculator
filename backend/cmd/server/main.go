package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebobo/investment_calculator/pkg/server"
	"github.com/jessevdk/go-flags"
)

var opt struct {
	GRPCAddr string `short:"g" long:"grpc-addr" default:":9092" description:"gRPC listen address"`
	HTTPAddr string `short:"h" long:"http-addr" default:":9090" description:"http listen address" required:"yes"`
	MSAddr   string `short:"m" long:"ms-addr" default:":9094" description:"save record micro service gRPC address"`
}

func main() {
	_, err := flags.ParseArgs(&opt, os.Args)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	server := server.New(server.Config{
		GRPCListenAddr: opt.GRPCAddr,
		HTTPListenAddr: opt.HTTPAddr,
	})

	e := server.Start()
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
