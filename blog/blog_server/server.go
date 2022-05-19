package main

import (
	"fmt"
	"grpc-masterclass/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

type server struct {
	// This is required. Otherwise won't compile
	blogpb.UnimplementedBlogServiceServer
}

func main() {
	// If there is a go crash - get file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed listeinig:", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("gRPC server starts listenting on port 50051...")
		if err := s.Serve(lis); err != nil {
			log.Fatalln("Filed to serve:", err)
		}
	}()

	// Wait for Control + C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until the signal is received
	<-ch
	s.Stop()    // Stopping the server
	lis.Close() // Closing the listener
	fmt.Println("\nServer stopped")
}
