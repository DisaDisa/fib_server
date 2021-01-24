package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	ps "github.com/DisaDisa/fib_server.git/grpcservice"
	garbler "github.com/michaelbironneau/garbler/lib"
	grpc "google.golang.org/grpc"
)

type Service struct{}

func (s *Service) Generate(ctx context.Context,
	req *ps.Request) (*ps.Response, error) {

	var err error
	response := new(ps.Response)

	requirements := garbler.MakeRequirements(req.Sample)
	response.Response, err = garbler.NewPassword(&requirements)

	return response, err
}

//CreateGRPCService creates gRPC server
func CreateGRPCService() {
	server := grpc.NewServer()

	instance := new(Service)

	ps.RegisterServiceServer(server, instance)

	listener, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
	fmt.Println("gRPC server is listening")
}
