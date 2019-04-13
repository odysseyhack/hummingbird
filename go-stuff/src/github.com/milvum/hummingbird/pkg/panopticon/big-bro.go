package panopticon

import (
	"log"
	"math"

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

func aggregateResults(messages []pb.MessageRequest, stream pb.HiberBridge_BatchStatusesClient) {
	// TODO super super dummy, no binning whatsoever!
	for _, message := range messages {
		loc := message.Info.Location
		intensity := urgencyToIntensity(message.Info.StatusMap)
		aggregate := pb.AggregateStatus{BinnedLocation: loc, CountedStatuses: &intensity}
		if err := stream.Send(&aggregate); err != nil {
			log.Printf("%v.Send(%v) = %v", stream, aggregate, err)
		}
	}
	_, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	return

}
