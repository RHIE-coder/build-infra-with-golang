package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "golang/external/grpc/middlebeat/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedMiddleBeatServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.GetSendMessage())

	if in.GetSendMessage() == "idiot" {
		log.Printf("client must be careful...")
	}

	if in.GetSendMessage() == "noob" {
		log.Fatalf("die, too angry")
		panic("#$#!$&%#!&")
	}

	return &pb.EchoResponse{EchoMessage: "ECHO: " + in.GetSendMessage()}, nil
}

func main() {
	port := 13335
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMiddleBeatServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		prefix := "[client]:"
		fmt.Println(prefix + "into go routine")
		fmt.Println(prefix + "wait 5 seconds")
		time.Sleep(5 * time.Second)
		fmt.Println(prefix + "waiting done")
		conn, err := grpc.Dial("localhost:13335", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		c := pb.NewMiddleBeatClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		r, err := c.Echo(ctx, &pb.EchoRequest{SendMessage: "hello~~~!"})
		if err != nil {
			panic(err)
		}

		fmt.Println(prefix + "Get MSG: " + r.GetEchoMessage())
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
