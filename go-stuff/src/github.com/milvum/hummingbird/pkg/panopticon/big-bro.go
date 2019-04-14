package panopticon

import (
	"log"
	"math"
	"google.golang.org/grpc"
	"context"
	"time"

	pb "github.com/milvum/hummingbird/proto"
)

func simpleSquaredDist(l1 pb.GeoLocation, l2 pb.GeoLocation) float64 {
	// We technically do not really need to take the square root either, as we just compare useless numbers
	return math.Pow((l1.Lon-l2.Lon), 2) + math.Pow((l1.Lat-l2.Lat), 2)
}

func urgencyToIntensity(urgency *pb.Statuses) (intensity pb.StatusIntensity) {
	intensity.WATER = pb.Urgency_value[urgency.WATER.String()]
	intensity.FIRST_AID = pb.Urgency_value[urgency.FIRST_AID.String()]
	intensity.FOOD = pb.Urgency_value[urgency.FOOD.String()]
	intensity.SHELTER = pb.Urgency_value[urgency.SHELTER.String()]
	intensity.PROTECTION = pb.Urgency_value[urgency.PROTECTION.String()]
	intensity.FIRE = pb.Urgency_value[urgency.FIRE.String()]
	intensity.ELECTRICITY = pb.Urgency_value[urgency.ELECTRICITY.String()]
	intensity.TOOLS = pb.Urgency_value[urgency.TOOLS.String()]
	intensity.TRANSPORT = pb.Urgency_value[urgency.TRANSPORT.String()]
	return
}

func aggregateResults(message pb.MessageRequest) pb.AggregateStatus {
	loc := message.Info.Location
	intensity := urgencyToIntensity(message.Info.StatusMap)
	return pb.AggregateStatus{BinnedLocation: loc, CountedStatuses: &intensity}
}

type hiberBridgeClient struct {
	destAddr string
	timeout  time.Duration
	opts     []grpc.DialOption
}

func newClient() *hiberBridgeClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	s := &hiberBridgeClient{opts: opts, timeout: 10 * time.Second}
	return s
}

func (c *hiberBridgeClient) uplink(messageRequest pb.MessageRequest) {
	conn, err := grpc.Dial(c.destAddr, c.opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewHiberBridgeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	hbp := aggregateResults(messageRequest)
	_, err = client.BatchStatuses(ctx, &hbp)
	if err != nil {
		log.Fatalf("No clue, something went wrong")
	}
}
