package main

import (
	"context"
	"fmt"
	"grpc-masterclass/calculator/calculatorpb"
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
