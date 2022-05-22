package main

import (
	"context"
	"fmt"
	"grpc-masterclass/blog/blogpb"
	"io"
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
	doUpdateBlog(c)
	doDeleteBlog(c)
	doListBlog(c)
}

func doListBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.ListBlogRequest{}
	resStream, err := c.ListBlog(context.Background(), req)
	if err != nil {
		log.Fatalln("error while calling GreetManyTimes RPC:", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			// reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalln("error while reading stream:", err)
		}
		fmt.Println(res)
	}

}

func doDeleteBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.DeleteBlogRequest{
		BlogId: "628a85f795a248097fc68a5d",
	}

	res, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling Blog RPC:", err)
	}

	fmt.Println(res)
}

func doUpdateBlog(c blogpb.BlogServiceClient) {
	req := &blogpb.Blog{
		Id:       "628a920244226b8a50c7fc24",
		AuthorId: "Which Agent?",
		Title:    "Not A Spy Blog",
		Content:  "This is an UPDATED record to the spy blog",
	}

	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: req})
	if err != nil {
		log.Fatalln("Error while calling Blog RPC:", err)
	}

	log.Println("Response: ", res)
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
