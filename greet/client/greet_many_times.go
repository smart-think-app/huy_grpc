package main

import (
	"context"
	"fmt"
	"huy_grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c proto.GreetServiceClient) {

	req := &proto.GreetRequest{}

	stream , err := c.GreetManyTimes(context.Background() , req)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for {
		msg , err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(msg.Result)
	}
}