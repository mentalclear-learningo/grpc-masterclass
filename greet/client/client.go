package main

import (
	"context"
	"fmt"
	"grpc-masterclass/greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Report: I'm client, connectin' to the server now...")

	certFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalln("failed loading certificates:", sslErr)
	}

	// cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // For insecure mode.
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
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

	// ClientStereaming
	doClientStreaming(c)

	// BiDirectional Streaming
	doBiDiStreaming(c)

	// Unary with Deadline
	doUnaryWithDeadline(c, 5000*time.Millisecond) // Should complete
	doUnaryWithDeadline(c, 1000*time.Millisecond) // Should timeout
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting Unary With Deadline RPC...")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dmitrii",
			LastName:  "Kilishek",
		},
	}

	// Context Timeout here for the client
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			log.Println("Error code:", statusErr.Code())
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("the deadline was exceeded")
			} else {
				log.Println("unexpected error:", statusErr)
			}
		} else {
			log.Fatalln("error while calling GreetWithDeadline RPC:", err)
		}
		return
	}

	log.Println("GreetWithDeadline Response: ", res.Result)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting BiDi Streaming RPC...")

	// Create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln("error while creating stream:", err)
		return
	}

	reqs := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Dmitrii",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Carmen",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Andrew",
			},
		},
	}

	waitc := make(chan struct{})
	// Send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range reqs {
			fmt.Println("Sending message:", req)
			stream.Send((*greetpb.GreetEveryoneRequest)(req))
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// Recieve a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("error while receiving:", err)
				break
			}
			fmt.Println("Received:", res.GetResult())
		}
		close(waitc)
	}()

	// Block until everythin is done
	<-waitc
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Client Streaming RPC...")

	reqs := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Dmitrii",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Carmen",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Andrew",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalln("error while calling LongGreet RPC:", err)
	}
	// Iterate over the slice and send each message
	for _, req := range reqs {
		fmt.Println("Sending request:", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("error while receiving response from LongGreet RPC:", err)
	}
	fmt.Println("LongGreet Response:", res)
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
