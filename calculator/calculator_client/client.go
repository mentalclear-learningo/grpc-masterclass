package main

import (
	"context"
	"fmt"
	"grpc-masterclass/calculator/calculatorpb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Report: I'm calculator client, connectin' to the server now...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("couldn't connect!", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	// Unary
	doUnary(c)

	// Server Streaming
	doServerStreaming(c)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Prime Dcomp Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12390392840,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalln("error while calling PrimeNumberDecomposition RPC:", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalln("error while reading stream:", err)
		}
		log.Println("Response from PrimeNumberDecomposition:", res.GetPrimeFactor())
	}
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling Sum RPC:", err)
	}

	log.Println("Response: ", res.SumResult)
}
