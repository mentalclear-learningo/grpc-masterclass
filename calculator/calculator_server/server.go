package main

import (
	"context"
	"fmt"
	"grpc-masterclass/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	// This is required, otherwise won't compile:
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Println("Recived Sum RPC:", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}

	return res, nil
}

func main() {
	fmt.Println("gRPC server starts listenting on port 50051...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed listeinig:", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Filed to serve:", err)
	}
}
