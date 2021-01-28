package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

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
	x, y := in.GetX(), in.GetY()
	if x > y {
		return &pb.Response{Response: nil}, errors.New("X must be less than Y")
	}
	response := make([]int32, 0, y-x+1)
	timeStart := time.Now()
	for i := x; i <= y; i++ {
		newVal, err := fibcalc.GetFibNimber(int(i))
		if err != nil {
			panic(err)
		}
		response = append(response, int32(newVal))
	}
	fmt.Println(int(time.Since(timeStart)))
	return &pb.Response{Response: response}, nil
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
