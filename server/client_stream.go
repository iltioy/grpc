package main

import (
	"io"
	"log"

	pb "github.com/iltioy/grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
        if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err!= nil {
            return err
        }
		log.Printf("got message: %s", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}