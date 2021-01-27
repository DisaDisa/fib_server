package grpcserver

import (
	"fmt"
	"log"
	"net"
	"sync"

	grpc "google.golang.org/grpc"
)

//CreateGRPCService creates gRPC server
func CreateGRPCService(wg *sync.WaitGroup) {
	defer wg.Done()

	server := grpc.NewServer()

	listener, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	fmt.Println("gRPC server is listening...")
	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
