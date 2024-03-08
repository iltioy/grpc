package main

import (
	"io"
	"log"

	pb "github.com/iltioy/grpc/proto"
)

func (s *helloServer) SayHelloBiderectionalStreaming(stream pb.GreetService_SayHelloBiderectionalStreamingServer) error {
	for {
        req, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }
		log.Printf("got message: %s", req.Name)
        if err := stream.Send(&pb.HelloResponse{Message: "Hello " + req.Name}); err != nil {
			return err
		}
    }
}