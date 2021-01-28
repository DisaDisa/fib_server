package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/DisaDisa/fib_server.git/grpcservice"
	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:8282"
)

func main() {
	connection, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()
	client := pb.NewGPRCServiceClient(connection)

	if len(os.Args) != 3 {
		fmt.Println("Pass index range fibonacci sequence")
		return
	}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("X parse error")
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Y parse error")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.FibGRPCHandler(ctx, &pb.Request{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatalf("Fib calc error: %v", err)
	}
	fmt.Println(r.Response)
}
