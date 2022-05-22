package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"huy_grpc/greet/proto"
	"log"
)

var address string = "0.0.0.0:50051"
func main() {
	conn , err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v" , err)
	}

	c := proto.NewGreetServiceClient(conn)
	//DoGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	doGreetEveryOne(c)
	defer conn.Close()
}
