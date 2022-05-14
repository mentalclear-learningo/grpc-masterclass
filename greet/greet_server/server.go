package main

import (
	"context"
	"fmt"
	"grpc-masterclass/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	// This is required. Otherwise won't compile
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet fucntion invoked with request:", req)
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
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
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Filed to serve:", err)
	}
}
