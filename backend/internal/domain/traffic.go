package domain

import "time"

type TrafficLevel string

const (
	TrafficLow    TrafficLevel = "low"
	TrafficMedium TrafficLevel = "medium"
	TrafficHigh   TrafficLevel = "high"
)

type TrafficObservation struct {
	EdgeID EdgeID `json:"edgeId"`
	//DensityVehKm float64 `json:"densityVehKm"`
	SpeedKmh float64 `json:"speedKmh"`
	Congestion float64 `json:"congestion"`
	ObservedAt time.Time `json:"observedAt"`
} 