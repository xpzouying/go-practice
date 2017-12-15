package main

import (
	"log"
	"net"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement Server1Server
type server struct{}

func (s *server) API1(ctx context.Context, st *Student) (*Response, error) {
	return &Response{
		Success: true,
		Message: "finish",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterServer1Server(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
