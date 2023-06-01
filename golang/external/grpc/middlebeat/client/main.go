package main

import (
	"context"
	pb "golang/external/grpc/middlebeat/protobuf"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	msg := os.Args[1]

	conn, err := grpc.Dial("localhost:13335", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMiddleBeatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	/* ECHO */
	for {
		r, err := c.Echo(ctx, &pb.EchoRequest{SendMessage: msg})
		if err != nil {
			log.Fatalf("could not echo: %v", err)
		}
		log.Printf("Get Echo Message: %s", r.GetEchoMessage())
		time.Sleep(1 * time.Second)
		/*
			2023/06/01 09:51:28 Get Echo Message: ECHO: hello
			2023/06/01 09:51:29 Get Echo Message: ECHO: hello
			2023/06/01 09:51:30 could not echo: rpc error: code = DeadlineExceeded desc = context deadline exceeded
			exit status 1
		*/
		/*
			noob
			2023/06/01 10:04:47 could not echo: rpc error: code = Unknown desc = panic triggered: #$#!$&%#!&
			exit status 1
		*/
	}

	/* STATUS */
	// r, err := c.Status(ctx, &pb.StatusRequest{Sender: "john"})
	// if err != nil {
	// 	log.Fatalf("could not get status: %v", err)
	// }
	// // fmt.Println(r.ProtoReflect().Descriptor())
	// jsonStr, _ := json.MarshalIndent(r, "", "  ")
	// log.Printf("Get Status From Server: %v", string(jsonStr))

}
