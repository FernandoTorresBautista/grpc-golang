package main

import (
	"context"
	"fmt"
	"grpc-course/hello/hellopb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)

	// helloUnary(c)
	// helloServerStreaming(c)
	goodbyeClientStreaming(c)
}

func helloUnary(c hellopb.HelloServiceClient) {
	fmt.Println("Starting unary RPC Hello")

	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "FirstName_fake",
			Prefix:    "Prefix_fake",
		},
	}

	res, err := c.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error, calling Hello RPC: \n%v", err)
	}

	log.Printf("Response Hello: %v", res.CustomHello)
}

func helloServerStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting server streaming RPC Hello")

	req := &hellopb.HelloManyLanguagesRequest{
		Hello: &hellopb.Hello{
			FirstName: "FirstName_fake",
			Prefix:    "Prefix_fake",
		},
	}

	restStream, err := c.HelloManyLanguages(context.Background(), req)

	if err != nil {
		log.Printf("Error calling Hello Many languages %v", err)
	}

	for {
		msg, err := restStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream %v", err)
		}
		log.Printf("Res from HML: %v\n", msg.GetHelloLanguage())
	}
}

func goodbyeClientStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting goodbye function")

	requests := []*hellopb.HellosGoodbyeRequest{
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "FN_fake_0",
				Prefix:    "00",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "FN_fake_1",
				Prefix:    "01",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "FN_fake_2",
				Prefix:    "02",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "FN_fake_3",
				Prefix:    "03",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "FN_fake_4",
				Prefix:    "04",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "FN_fake_5",
				Prefix:    "05",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "FN_fake_6",
				Prefix:    "06",
			},
		},
	}

	stream, err := c.HellosGoodbye(context.Background())

	if err != nil {
		log.Printf("Error calling goodbye %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	goodbye, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Error goodbye receive %v", err)
	}

	fmt.Println("Response server for client stream: ", goodbye)
}
