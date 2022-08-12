package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"grpc-course/hello/hellopb"

	"google.golang.org/grpc"
)

type server struct{}

// hello service unary
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

// hello server streaming
func (*server) HelloManyLanguages(req *hellopb.HelloManyLanguagesRequest, stream hellopb.HelloService_HelloManyLanguagesServer) error {
	fmt.Printf("Hello Many times function was invoked with %v\n", req)

	langs := [10]string{"Salut! ", "Hello! ", "Ni hao! ", "Alô! ", "Privyét! ", "Schalom! ", "Hola ! ", "Yassou! ", "Hej! ", "Konnichiwa! "}

	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	for _, helloLang := range langs {
		helloLanguage := helloLang + prefix + " " + firstName

		res := &hellopb.HelloManyLanguagesResponse{
			HelloLanguage: helloLanguage,
		}
		// send the response for the actual language
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond) // one second
	}

	return nil
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
