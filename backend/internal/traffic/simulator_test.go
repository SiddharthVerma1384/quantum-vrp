package traffic

import (
	"testing"
	"time"

	"quantum-vrp/backend/internal/domain"
)

func TestSimulatorProducesTrafficObservation(t *testing.T) {
	start := time.Date(
		2026,
		1,
		1,
		10,
		0,
		0,
		0,
		time.UTC,
	)

	simulator := NewSimulator(start)

	edge := domain.Edge{
		ID:           "E01",
		DistanceKm:   5,
		BaseSpeedKmh: 50,
	}

	observation := simulator.NextObservation(edge)

	if observation.EdgeID != "E01" {
		t.Fatalf(
			"expected edge E01, got %s",
			observation.EdgeID,
		)
	}

	if observation.SpeedKmh <= 0 {
		t.Fatalf(
			"expected positive speed, got %.2f",
			observation.SpeedKmh,
		)
	}

	if observation.Congestion < 0 ||
		observation.Congestion > 1 {
		t.Fatalf(
			"invalid congestion %.2f",
			observation.Congestion,
		)
	}

	if observation.ObservedAt != start {
		t.Fatalf(
			"expected observation time %v, got %v",
			start,
			observation.ObservedAt,
		)
	}
}