package main

import (
	"context"
	"io"
	"log"

	pb "github.com/iltioy/grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err!= nil {
        log.Fatalf("could not save names: %v", err)
    }

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", names)
		}
		log.Printf("got message: %v", message)
	}
	log.Printf("streaming finished")
}