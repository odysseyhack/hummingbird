package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/milvum/hummingbird/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:8081", "The server address")
)

func cast(client pb.HummingbirdClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.CastStatuses(ctx, &pb.StatusRequest{Origin: "wtf"})
	if err != nil {
		log.Fatalf("No clue, something went wrong")
	}
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHummingbirdClient(conn)
	cast(client)

	fmt.Println("hello world")
}
