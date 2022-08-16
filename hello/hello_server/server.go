package main

import (
	"context"
	"fmt"
	"io"
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

// hello client streaming
func (*server) HellosGoodbye(stream hellopb.HelloService_HellosGoodbyeServer) error {
	fmt.Println("Goodbye function was invoked")

	goodbye := "Goodbye guys: "

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// Once is finished the stream we gonna send the response
			return stream.SendAndClose(&hellopb.HellosGoodbyeResponse{
				Goodbye: goodbye,
			})
		}

		if err != nil {
			log.Fatalf("Error reading the client stream %v", err)
		}

		firstName := req.GetHello().GetFirstName()
		prefix := req.GetHello().GetPrefix()

		goodbye += prefix + " " + firstName + " "
	}
}

// goodbye bidirectional function
func (*server) Goodbye(stream hellopb.HelloService_GoodbyeServer) error {
	fmt.Println("Goodbye bidirectional function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error reading the client stream %v", err)
			return nil
		}

		firstName := req.GetHello().GetFirstName()
		prefix := req.GetHello().GetPrefix()

		goodbye := "Goodbye " + prefix + " " + firstName + " :("

		sendErr := stream.Send(&hellopb.GoodbyeResponse{
			Goodbye: goodbye,
		})

		if sendErr != nil {
			log.Fatalf("Error sending to the client %v", sendErr)
		}
	}
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
