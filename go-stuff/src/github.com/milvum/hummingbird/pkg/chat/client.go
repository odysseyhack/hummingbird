package chat

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"

	pb "github.com/milvum/hummingbird/proto"
)

// TODO: We should have a bidirectional connection instead of this approach
type hummingbirdClient struct {
	destAddr string
	timeout  time.Duration
	opts     []grpc.DialOption
}

func newClient() *hummingbirdClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	s := &hummingbirdClient{opts: opts, timeout: 10 * time.Second}
	return s
}

func (c *hummingbirdClient) BroadcastMessage(messageRequest pb.MessageRequest) {
	conn, err := grpc.Dial(c.destAddr, c.opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewHummingbirdClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err = client.SendMessage(ctx, &messageRequest)
	if err != nil {
		log.Fatalf("No clue, something went wrong")
	}
}

func (c *hummingbirdClient) BroadcastStatus(statusRequest pb.StatusRequest) {
	conn, err := grpc.Dial(c.destAddr, c.opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewHummingbirdClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err = client.CastStatuses(ctx, &statusRequest)
	if err != nil {
		log.Fatalf("No clue, something went wrong")
	}
}
