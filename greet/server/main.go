package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"huy_grpc/greet/proto"
	"log"
	"net"
)

var address string = "0.0.0.0:50051"

type Server struct{
	proto.GreetServiceServer
}

func main() {
	lis , err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Fail with %v\n" , err)
	}

	log.Printf("Listening on %v" , address)

	s := grpc.NewServer()

	proto.RegisterGreetServiceServer(s,&Server{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Fail with %v" , err)
	}
}
