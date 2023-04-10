package main

import (
	"gprc/protos/currency"
	"gprc/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	c := server.NewCurrency(log)

	// register the currency server
	currency.RegisterCurrencyServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	// start server
	gs.Serve(l)
}
