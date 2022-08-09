package main

import (
	"fmt"
	"grpc-course/hello/hellopb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running")
	cc, err := grpc.Dial("locahost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connet %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)
	// implementar el servicio unitario
}
