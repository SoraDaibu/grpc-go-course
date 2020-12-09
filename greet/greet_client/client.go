package main

import (
	"fmt"
	"log"

	"github.com/SoraDaibu/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
