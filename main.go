package main

import (
	"sync"

	"github.com/DisaDisa/fib_server.git/grpcserver"
	"github.com/DisaDisa/fib_server.git/httpserver"
)

func main() {
	ServersWG := &sync.WaitGroup{}
	ServersWG.Add(2)
	go httpserver.CreateServer(ServersWG)
	go grpcserver.CreateGRPCService(ServersWG)
	ServersWG.Wait()
}
