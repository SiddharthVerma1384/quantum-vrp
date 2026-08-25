package graph

import (
	"fmt"

	"quantum-vrp/backend/internal/domain"
)

func ApplyTrafficObservation(
	graph *domain.Graph,
	observation domain.TrafficObservation,
) error {

	edge, exists := graph.Edges[observation.EdgeID]

	if !exists {
		return fmt.Errorf(
			"edge %s not found in graph",
			observation.EdgeID,
		)
	}

	if observation.SpeedKmh <= 0 {
		return fmt.Errorf(
			"invalid speed %.2f for edge %s",
			observation.SpeedKmh,
			observation.EdgeID,
		)
	}

	edge.CurrentSpeedKmh = observation.SpeedKmh

	edge.Congestion = observation.Congestion

	edge.TravelTimeMin =
		(edge.DistanceKm / observation.SpeedKmh) * 60

	graph.Edges[observation.EdgeID] = edge

	return nil
}