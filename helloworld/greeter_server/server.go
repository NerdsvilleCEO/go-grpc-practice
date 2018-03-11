package main

import (
	"net"

	"google.golang.org/grpc"
)

func ListenRPC() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		return nil, err
	}
	server := grpc.NewServer()
	return server, server.Serve(lis)
}

func main() {
	ListenRPC()
	for {

	}
}
