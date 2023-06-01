package main

/*
 - grpc dial method
 - 채널 + Context
 - wg + mutex
 - grpc middleware
 - container
 - net IPNet
 - beanstalk
*/
import (
	"context"
	"fmt"
	pb "golang/external/grpc/heartbeat/protobuf"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHerBeaServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.GetSendMessage())
	log.Debug(in.GetSendMessage())
	if in.GetSendMessage() == "idiot" {
		log.Warnf("client must be careful...")
	}

	if in.GetSendMessage() == "noob" {
		log.Errorf("die, too angry")
		panic("#$#!$&%#!&")
	}
	return &pb.EchoResponse{EchoMessage: "ECHO: " + in.GetSendMessage()}, nil
}

func (s *server) Status(ctx context.Context, in *pb.StatusRequest) (*pb.StatusResponse, error) {
	log.Printf("Some Client who [%s] Wants to Check", in.GetSender())

	members := []*pb.Member{
		{
			Name: "alice",
			Age:  uint32(12),
			Role: "anony",
		},
		{
			Name: "bob",
			Age:  uint32(24),
			Role: "vip",
		},
		{
			Name: "charlie",
			Age:  uint32(53),
			Role: "admin",
		},
	}

	return &pb.StatusResponse{
		Receiver:       "rhie",
		Status:         true,
		Message:        "Yes, I'm alive",
		AllowedMembers: members,
		Langs:          []string{"python", "java", "golang"},
	}, nil
}

func main() {
	port := 13335
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHerBeaServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
