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

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		panic("hello")
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
	log.ErrorKey = "grpc.error"
	logrusEntry := log.NewEntry(log.StandardLogger())
	// Define customfunc to handle panic
	customFunc := func(p interface{}) (err error) {
		fmt.Println(p)
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),

			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	pb.RegisterHerBeaServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
