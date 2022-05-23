#!/bin/bash

# Outdated:
#protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet/greetpb/greet.proto

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    calculator/calculatorpb/calculator.proto

# Generate Go code 
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
blog/blogpb/blog.proto

# Generate Python code 
protoc -I=blog/blogpb --python_out=blog/blogpb blog/blogpb/blog.proto
