package main

import (
	"huy_grpc/greet/proto"
	"io"
	"log"
)

func(s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {

	res := ""
	for {
		req , err := stream.Recv()
		if err ==io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{Result: res})
		}

		if err  != nil {
			log.Fatalf(err.Error())
		}

		res += req.FirstName + " "
	}

	return nil
}
