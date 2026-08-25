package domain

type NodeID string
type EdgeID string

type Node struct {
	ID        NodeID `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Edge struct {
	ID   EdgeID `json:"id"`
	From NodeID `json:"from"`
	To   NodeID `json:"to"`

	DistanceKm float64 `json:"distanceKm"`

	// Normal/free-flow road speed.
	BaseSpeedKmh float64 `json:"baseSpeedKmh"`

	// Normal/free-flow travel time.
	BaseTravelTimeMin float64 `json:"baseTravelTimeMin"`

	// Current observed speed under traffic.
	CurrentSpeedKmh float64 `json:"currentSpeedKmh"`

	// Dynamic travel time after traffic is applied.
	TravelTimeMin float64 `json:"travelTimeMin"`

	// 0 = free-flow, 1 = extremely congested.
	Congestion float64 `json:"congestion"`
}

type Graph struct {
	Nodes map[NodeID]Node `json:"nodes"`
	Edges map[EdgeID]Edge `json:"edges"`
}