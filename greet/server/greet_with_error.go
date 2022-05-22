package main

import (
	"context"
	"errors"
	"huy_grpc/greet/proto"
)

func(s *Server) GreetWithErr(ctx context.Context , req *proto.GreetRequest) (*proto.GreetResponse , error) {
	return nil , errors.New("This is Error")
}
