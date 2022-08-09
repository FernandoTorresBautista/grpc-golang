package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc-course/hello/hellopb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Printf("Hello function was invoked with %v\n", req)

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	customHello := "Welcome ! " + prefix + " " + firstName

	res := &hellopb.HelloResponse{
		CustomHello: customHello,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello, Go Server is running")

	// escuchar tcp, y puerto por default de grpc
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
