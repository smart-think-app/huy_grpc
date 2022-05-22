package main

import (
	"context"
	"fmt"
	"huy_grpc/greet/proto"
	"log"
)

func DoGreet(c proto.GreetServiceClient) {
	res , err := c.Greet(context.Background() , &proto.GreetRequest{FirstName: "Huy"})
	if err != nil{
		log.Fatalf(err.Error())
	}
	fmt.Println(res.Result)
}