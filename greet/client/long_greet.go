package main

import (
	"context"
	"fmt"
	"huy_grpc/greet/proto"
	"log"
)

func doLongGreet(client proto.GreetServiceClient) {
	reqs := []*proto.GreetRequest{
		{FirstName: "Huy"},
		{FirstName: "Huy1"},
	}

	stream , err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _ , req := range reqs {
		stream.Send(req)
	}

	res , err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(res)
}
