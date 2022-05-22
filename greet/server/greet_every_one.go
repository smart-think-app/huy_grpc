package main

import (
	"huy_grpc/greet/proto"
	"io"
	"log"
)

func(s *Server) GreetEveryOne(stream proto.GreetService_GreetEveryOneServer) error {
	for {
		req , err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf(err.Error())
		}

		res := "Hello " + req.FirstName + "!"

		errSend := stream.Send(&proto.GreetResponse{
			Result: res,
		})

		if errSend != nil {
			log.Fatalf(errSend.Error())
		}


	}
}
