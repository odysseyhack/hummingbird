package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "../../go-stuff/src/github.com/milvum/hummingbird/proto"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:8081", "The server address")
)

type HiberBridgeClient interface {
	BatchStatuses(ctx context.Context, opts ...grpc.CallOption) (HiberBridge_BatchStatusesClient, error)
}

type hiberBridgeClient struct {
	cc *grpc.ClientConn
}

func NewHiberBridgeClient(cc *grpc.ClientConn) HiberBridgeClient {
	return &hiberBridgeClient{cc}
}

func (c *hiberBridgeClient) BatchStatuses(ctx context.Context, opts ...grpc.CallOption) (pb.HiberBridge_BatchStatusesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HiberBridge_serviceDesc.Streams[0], "/hummingbird.proto.HiberBridge/BatchStatuses", opts...)
	if err := stream.Send(point); err != nil {
		if err == io.EOF {
			break
		}
}

func newClient() *HiberBridgeClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	s := &HiberBridgeClient{opts: opts, timeout: 10 * time.Second}
	return s
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
	client := pb.NewHiberBridgeClient(conn)
	client.BatchStatuses(context.Background())
}
