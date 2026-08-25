package main

import (
	"fmt"
	"time"

	"quantum-vrp/backend/internal/domain"
	"quantum-vrp/backend/internal/traffic"
)

func main() {
	startTime := time.Now()

	simulator := traffic.NewSimulator(startTime)

	edge := domain.Edge{
		ID:           "E01",
		DistanceKm:   5,
		BaseSpeedKmh: 50,
	}

	fmt.Println("========================================")
	fmt.Println("TRAFFIC SIMULATION")
	fmt.Println("========================================")

	fmt.Printf(
		"%-10s %-10s %-15s %-12s\n",
		"Time",
		"Edge",
		"Speed",
		"Congestion",
	)

	fmt.Println("----------------------------------------")

	for i := 0; i < 10; i++ {
		observation := simulator.NextObservation(edge)

		fmt.Printf(
			"%-10s %-10s %-15.2f %-12.2f\n",
			observation.ObservedAt.Format("15:04"),
			observation.EdgeID,
			observation.SpeedKmh,
			observation.Congestion,
		)
	}
}