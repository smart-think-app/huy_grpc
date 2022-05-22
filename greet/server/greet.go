package main

import (
	"context"
	"huy_grpc/greet/proto"
)

func(s *Server) Greet(ctx context.Context , request *proto.GreetRequest) (*proto.GreetResponse , error) {
	return &proto.GreetResponse{Result: "Hello" + request.FirstName} , nil
}