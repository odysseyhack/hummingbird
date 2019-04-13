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
	/*intensity.WATER = urgency.WATER
	intensity.FIRST_AID = urgency.FIRST_AID
	intensity.FOOD = urgency.FOOD
	intensity.SHELTER = urgency.SHELTER
	intensity.PROTECTION = urgency.PROTECTION
	intensity.FIRE = urgency.FIRE
	intensity.ELECTRICITY = urgency.ELECTRICITY
	intensity.TOOLS = urgency.TOOLS
	intensity.TRANSPORT = urgency.TRANSPORT */
	// TODO: Fix :
	// cannot use urgency.TRANSPORT (type hummingbird_proto.Urgency) as type int32 in assignment
	intensity.FIRE = 10
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
