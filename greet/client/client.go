package main

import (
	"context"
	"fmt"
	"grpc-masterclass/greet/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Report: I'm client, connectin' to the server now...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("couldn't connect!", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	// Unary
	doUnary(c)

	// Server Stream
	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Server Streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dmitrii",
			LastName:  "Kilishek",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling GreetManyTimes RPC:", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalln("error while reading stream:", err)
		}
		log.Println("Response from GreetManyTimes:", msg.GetResult())
	}
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dmitrii",
			LastName:  "Kilishek",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling Greet RPC:", err)
	}

	log.Println("Response: ", res.Result)
}
