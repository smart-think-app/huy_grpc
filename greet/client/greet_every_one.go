package main

import (
	"context"
	"fmt"
	"huy_grpc/greet/proto"
	"io"
	"log"
)

func doGreetEveryOne(c proto.GreetServiceClient){
	stream , err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Huy"},
		{FirstName: "Tran"},
	}

	waitc := make(chan struct{})

	go func() {
		for _ , req := range reqs {
			fmt.Println(req)
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res , err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Println(res)
		}
		close(waitc)
	}()

	<-waitc
}