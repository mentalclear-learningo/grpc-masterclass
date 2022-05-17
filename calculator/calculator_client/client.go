package main

import (
	"context"
	"fmt"
	"grpc-masterclass/calculator/calculatorpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

	// Client Streaming
	doClientStreaming(c)

	// BiDi Streaming
	doBiDiStreaming(c)

	// Error Unary
	doErrorUnary(c)
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Error Unary RPC request...")

	numbers := []int32{1, 4, 9, 16, 25, -9, -1}
	for _, num := range numbers {
		doErrorCall(c, num)
	}
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, number int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: number})
	if err != nil {

		respErr, ok := status.FromError(err)
		if ok {
			log.Println("Error message from the server:", respErr.Message())
			log.Println("Error code:", respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				log.Println("Probably a negative number was sent!")
				return
			}
		} else {
			log.Fatalln("Big Error calling SquareRoot RPC:", err)
			return
		}
	}
	log.Printf("Result for Square Root of number: %v, equals: %v\n", number, res.GetNumberRoot())
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting FindMaximum BiDi Streaming RPC...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalln("error while opening stream:", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		nums := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, num := range nums {
			fmt.Println("Sending number:", num)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: num,
			})
			time.Sleep(300 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("error while reading server stream:", err)
				break
			}
			fmt.Println("Received a new maximum:", res.GetMaximum())
		}
		close(waitc)
	}()
	<-waitc
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting ComputeAverage Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalln("error while opening stream:", err)
	}

	nums := []int32{3, 5, 9, 54, 23}
	for _, num := range nums {
		log.Println("Sending number:", num)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: num,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("error while while receiving response", err)
	}
	fmt.Println("The average is:", res.GetAverage())
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
