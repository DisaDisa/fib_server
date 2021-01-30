package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/DisaDisa/fib_server.git/grpcservice"
	"github.com/DisaDisa/fib_server.git/server/fibcalc"
	grpc "google.golang.org/grpc"
)

const (
	port = ":8282"
)

type fibServer struct {
	pb.UnimplementedGPRCServiceServer
}

func (s *fibServer) FibGRPCHandler(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	x, y := int(in.GetX()), int(in.GetY())
	if x > y {
		return &pb.Response{Response: nil}, errors.New("X must be less than Y")
	}
	if x <= 0 {
		return &pb.Response{Response: nil}, errors.New("X must be greater than 0")
	}
	if y <= 0 {
		return &pb.Response{Response: nil}, errors.New("Y must be greater than 0")
	}

	response := fibcalc.GetFibRange(x, y)
	response32 := make([]int32, 0, len(response))
	for _, v := range response {
		response32 = append(response32, int32(v))
	}
	return &pb.Response{Response: response32}, nil
}

//CreateGRPCService creates gRPC server
func CreateGRPCService(wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	server := grpc.NewServer()
	pb.RegisterGPRCServiceServer(server, &fibServer{})

	fmt.Println("gRPC server is listening...")
	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
