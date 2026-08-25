package traffic

import (
	"math"
	"time"

	"quantum-vrp/backend/internal/domain"
)

type Simulator struct {
	StartTime time.Time
	Step      int
}

func NewSimulator(startTime time.Time) *Simulator {
	return &Simulator{
		StartTime: startTime,
	}
}

func (s *Simulator) NextObservation(
	edge domain.Edge,
) domain.TrafficObservation {

	currentTime := s.StartTime.Add(
		time.Duration(s.Step) * time.Minute,
	)

	// Creates a smooth traffic pattern between
	// approximately 0.2 and 0.8 congestion.
	phase := float64(s.Step) / 5.0

	congestion :=
		0.5 +
			0.3*math.Sin(phase)

	speedFactor :=
		1.0 - congestion*0.8

	speed := edge.BaseSpeedKmh * speedFactor

	if speed < 5 {
		speed = 5
	}

	s.Step++

	return domain.TrafficObservation{
		EdgeID:      edge.ID,
		SpeedKmh:    speed,
		Congestion:  congestion,
		ObservedAt:  currentTime,
	}
}