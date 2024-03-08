package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/iltioy/grpc/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("bidirectional streaming started")
	stream, err := client.SayHelloBiderectionalStreaming(context.Background())
	if err!= nil {
        log.Fatalf("could not create stream: %v", err)
    }
	waitc := make(chan struct{})
	go func() {
		for {
			messsage, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err!= nil {
                log.Fatalf("could not receive message: %v", err)
            }
			log.Println(messsage)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		err := stream.Send(&pb.HelloRequest{Name: name})
        if err!= nil {
            log.Fatalf("could not send message: %v", err)
        }
		time.Sleep(1 * time.Second)
	}
	log.Printf("bidirectional streaming finished")
}