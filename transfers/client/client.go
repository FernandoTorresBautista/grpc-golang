package main

import (
	"context"
	"fmt"
	"grpc-course/transfers/transferspb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := transferspb.NewTransfersServiceClient(cc)

	transfer(c)
}

func transfer(c transferspb.TransfersServiceClient) {
	fmt.Println("Starting unary RPC Transfer")

	// // error account
	// req := &transferspb.TransfersRequest{
	// 	OriginAccount:   "87655678909877",
	// 	ReceiverAccount: "98555682009402",
	// 	Amount:          100,
	// }
	// error amount
	// req := &transferspb.TransfersRequest{
	// 	OriginAccount:   "87655678909876",
	// 	ReceiverAccount: "98555682009402",
	// 	Amount:          0,
	// }
	// valid request
	req := &transferspb.TransfersRequest{
		OriginAccount:   "87655678909876",
		ReceiverAccount: "98555682009402",
		Amount:          100,
	}

	res, err := c.Transfer(context.Background(), req)

	if err != nil {

		respErr, exception := status.FromError(err)

		if exception {
			fmt.Printf("---Exception message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("Sorry the amount has to be more than 0")
				return
			}
			if respErr.Code() == codes.NotFound {
				fmt.Println("It looks like you mistyped the source account number")
				return
			}
		} else {
			log.Fatalf("Fatal Error calling Transfer: %v", err)
			return
		}
	}

	log.Printf("Response Transfer: %v", res.GetOperationDate())

}
