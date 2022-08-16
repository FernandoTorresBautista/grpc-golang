package main

import (
	"context"
	"fmt"
	"grpc-course/transfers/transferspb"
	"log"
	"net"
	"time"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

// the server choose when how many responses send for the cient

func (*server) Transfer(ctx context.Context, req *transferspb.TransfersRequest) (*transferspb.TransfersResponse, error) {
	fmt.Println("Starting transfer", req)

	originAccount := req.GetOriginAccount()
	//receiverAccount := req.GetReceiverAccount()

	amount := req.GetAmount()

	// invalid amount
	if amount <= 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Error amount less or equal to 0 with amount $ %v", amount),
		)
	}

	// example (hardcode) for an accont not valid
	if originAccount != "87655678909876" {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("This account does not belong to the bank: %v", originAccount),
		)
	}

	// logic to validate accounts number
	// transfer logic code

	year, month, day := time.Now().Date()

	res := &transferspb.TransfersResponse{
		OperationDate: &date.Date{
			Year:  int32(year),
			Month: int32(month),
			Day:   int32(day),
		},
	}

	return res, nil
}

func main() {
	fmt.Println("Transfers, Go Server is running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	transferspb.RegisterTransfersServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
