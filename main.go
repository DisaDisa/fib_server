package main

import (
	"github.com/DisaDisa/fib_server.git/grpcserver"
	"github.com/DisaDisa/fib_server.git/httpserver"
)

func main() {
	httpserver.CreateServer()
	grpcserver.CreateGRPCService()
}
