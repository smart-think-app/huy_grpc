package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"huy_grpc/greet/proto"
	"log"
)

func doError(c proto.GreetServiceClient) {
	res , err := c.GreetWithErr(context.Background() , &proto.GreetRequest{})
	if err != nil {
		e , ok := status.FromError(err)
		if ok {
			log.Fatalf("GrpC Error Message %s" , e.Message())
		} else {
			log.Fatalf("Non GrpC Error Message %s" , err.Error())
		}
	}

	fmt.Println(res)
}
