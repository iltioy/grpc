package main

import (
	"context"
	"log"
	"time"

	pb "github.com/iltioy/grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
        log.Fatalf("could not send names: %v", err)
    }

	for _, name := range names.Names {
        if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("error while sending: %v", err)
		}
		log.Printf("send the request with name: %s", name)
		time.Sleep(2 * time.Second)
    }

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err!= nil {
        log.Fatalf("could not receive response: %v", err)
    }
	log.Printf("%v", res.Messages)
}