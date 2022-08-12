package main

import (
	"context"
	"fmt"
	"grpc-course/hello/hellopb"
	"io"
	"log"

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
	helloServerStreaming(c)
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
