package graph

import (
	"testing"

	"quantum-vrp/backend/internal/domain"
)

func TestApplyTrafficObservation(t *testing.T) {
	graph := domain.Graph{
		Nodes: map[domain.NodeID]domain.Node{
			"N01": {
				ID: "N01",
			},
			"N02": {
				ID: "N02",
			},
		},

		Edges: map[domain.EdgeID]domain.Edge{
			"E01": {
				ID:                 "E01",
				From:               "N01",
				To:                 "N02",
				DistanceKm:         5,
				BaseTravelTimeMin:  6,
				TravelTimeMin:      6,
				Congestion:         0,
			},
		},
	}

	observation := domain.TrafficObservation{
		EdgeID:      "E01",
		SpeedKmh:    20,
		Congestion:  0.8,
	}

	err := ApplyTrafficObservation(
		&graph,
		observation,
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	edge := graph.Edges["E01"]

	expectedTravelTime := 15.0

	if edge.TravelTimeMin != expectedTravelTime {
		t.Fatalf(
			"expected travel time %.2f, got %.2f",
			expectedTravelTime,
			edge.TravelTimeMin,
		)
	}

	if edge.Congestion != 0.8 {
		t.Fatalf(
			"expected congestion %.2f, got %.2f",
			0.8,
			edge.Congestion,
		)
	}
}