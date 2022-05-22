package main

import (
	"context"
	"fmt"
	"grpc-masterclass/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type server struct {
	// This is required. Otherwise won't compile
	blogpb.UnimplementedBlogServiceServer
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := req.GetBlog()
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("internal error inserting into the database", err))
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("internal error converting ObjectID", err))
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(), // ObjectID must be Hex
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	data := &blogItem{}                           // A data structure that will accept the results from collection findOne req from mongo
	blogId := req.GetBlogId()                     // Getting Blog's ObjectID from the request
	oid, err := primitive.ObjectIDFromHex(blogId) // Get ObjectID from hex
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("cannot parse provided ID %v", err))
	}

	// Based on the code examples from Go mongo package:
	errDecode := collection.FindOne(ctx, bson.M{"_id": oid}).Decode(data) // Decode puts data into empty data struct
	if errDecode != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection.
		if errDecode == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("cannot find the blog with the ID specified: %v", errDecode))
		}
		log.Fatal(errDecode)
	}

	// result := collection.FindOne(ctx, bson.M{"_id": oid}) // Filter: bson.M{"_id": oid} here is the bson representation of ObjectID

	// // Decoding result into the data structure
	// if err := result.Decode(data); err != nil {
	// 	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("cannot find the blog with the ID specified: %v", err))
	// }

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(), // ObjectID must be Hex
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}

func (*server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	data := &blogItem{}
	blog := req.GetBlog()

	oid, err := primitive.ObjectIDFromHex(blog.GetId()) // Get ObjectID from hex
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("cannot parse provided ID %v", err))
	}

	filter := bson.M{"_id": oid}
	errDecode := collection.FindOne(ctx, filter).Decode(data)
	if errDecode != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection.
		if errDecode == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("cannot find the blog with the ID specified: %v", errDecode))
		}
		log.Fatal(errDecode)
	}

	// Update internal struct
	data.AuthorID = blog.GetAuthorId()
	data.Title = blog.GetTitle()
	data.Content = blog.GetContent()

	_, updErr := collection.ReplaceOne(ctx, filter, data)
	if updErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot update object in MongoDB: %v", updErr),
		)
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog{
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}

func (*server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetBlogId()) // Get ObjectID from hex
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("cannot parse provided ID %v", err))
	}

	delRes, delErr := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if delErr != nil {
		return nil, fmt.Errorf("error deleting record %v", delErr)
	}

	return &blogpb.DeleteBlogResponse{
		Deleted: fmt.Sprintf("%v", delRes.DeletedCount),
	}, nil
}

func (*server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknow error in MongoDB: %v", err),
		)
	}
	defer cursor.Close(context.Background())

	results := []blogItem{}
	if err = cursor.All(context.Background(), &results); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("error decoding data from MongoDB: %v", err),
		)
	}

	for _, result := range results {
		res := &blogpb.ListBlogResponse{
			Blog: &blogpb.Blog{
				Id:       result.ID.Hex(),
				AuthorId: result.AuthorID,
				Title:    result.Title,
				Content:  result.Content,
			},
		}
		stream.Send(res)
	}

	return nil
}

func main() {
	// If there is a go crash - get file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Connect to mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln("Error connecting to the DB", err)
	}

	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Failed listeinig:", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

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
	//close mongodb connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	fmt.Println("\nServer gracefully stopped")
}
