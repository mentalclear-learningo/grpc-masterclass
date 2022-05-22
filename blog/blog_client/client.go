package main

import (
	"context"
	"fmt"
	"grpc-masterclass/blog/blogpb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Report: I'm a Blog client, connectin' to the server now...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // For insecure mode.
	if err != nil {
		log.Fatalln("couldn't connect!", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	doCreateBlog(c)
	doReadBlog(c)
}

func doReadBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.ReadBlogRequest{
		BlogId: "6286f2506428f4c313163592",
		// BlogId: "6286f2506428f4c313163591",
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling Blog RPC:", err)
	}

	log.Println("Response: ", res)
}

func doCreateBlog(c blogpb.BlogServiceClient) {
	blog := &blogpb.Blog{
		AuthorId: "Agent 007",
		Title:    "Spy Blog",
		Content:  "This is a new record to the spy blog",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalln("unexpected error creating a blog", err)
	}
	fmt.Println("The blog have been created!", res)
}
