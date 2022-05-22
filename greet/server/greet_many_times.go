package main

import (
	"huy_grpc/greet/proto"
	"strconv"
)

func(s *Server) GreetManyTimes(request *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	for i := 0 ; i < 10;i++ {
		stream.Send(&proto.GreetResponse{Result: strconv.Itoa(i)})
	}
	return nil
}
