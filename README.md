# grpc-masterclass
Working on gRPC [Golang] Master Class: Build Modern API and Microservice from Packt

Link in O'Reilly: https://learning-oreilly-com.lcpl.idm.oclc.org/videos/grpc-golang-master/9781838555467/9781838555467-video1_1/

The course is a bit outdated, but with slight changes it works good enough.

### Setup plugins

`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

### Local SSL/TLS Setup  

Make sure to run instructions.sh to setup local certificates after cloning the repo.  
This is required for the local server to work with SSL/TLS.  

