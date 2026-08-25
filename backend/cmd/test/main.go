package main

import (
	"encoding/json"
	"fmt"
	"time"

	"quantum-vrp/backend/internal/domain"
)

func main() {

	// -----------------------------------
	// 1. Create some nodes
	// -----------------------------------

	nodes := map[domain.NodeID]domain.Node{
		"N01": {
			ID:        "N01",
			Latitude:  12.9716,
			Longitude: 77.5946,
		},
		"N02": {
			ID:        "N02",
			Latitude:  12.9750,
			Longitude: 77.6000,
		},
		"N03": {
			ID:        "N03",
			Latitude:  12.9800,
			Longitude: 77.6050,
		},
	}

	// -----------------------------------
	// 2. Create roads
	// -----------------------------------

	edges := map[domain.EdgeID]domain.Edge{

		"E01": {
			ID:                 "E01",
			From:               "N01",
			To:                 "N02",
			DistanceKm:         5,
			BaseTravelTimeMin: 6,
			TravelTimeMin:      6,
			Congestion:         0,
		},

		"E02": {
			ID:                 "E02",
			From:               "N02",
			To:                 "N03",
			DistanceKm:         4,
			BaseTravelTimeMin: 5,
			TravelTimeMin:      5,
			Congestion:         0,
		},
	}

	graph := domain.Graph{
		Nodes: nodes,
		Edges: edges,
	}

	// -----------------------------------
	// 3. Display initial graph
	// -----------------------------------

	fmt.Println("================================")
	fmt.Println("INITIAL NETWORK")
	fmt.Println("================================")

	printJSON(graph)

	// -----------------------------------
	// 4. Simulate traffic observation
	// -----------------------------------

	traffic := domain.TrafficObservation{
		EdgeID:      "E01",
		SpeedKmh:    18,
		Congestion:  0.82,
		ObservedAt:  time.Now(),
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("TRAFFIC OBSERVATION")
	fmt.Println("================================")

	printJSON(traffic)

	// -----------------------------------
	// 5. Show what the traffic means
	// -----------------------------------

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("INTERPRETATION")
	fmt.Println("================================")

	fmt.Printf(
		"Edge %s currently has %.0f%% congestion\n",
		traffic.EdgeID,
		traffic.Congestion*100,
	)

	fmt.Printf(
		"Current speed: %.1f km/h\n",
		traffic.SpeedKmh,
	)
}

func printJSON(value any) {

	data, err := json.MarshalIndent(
		value,
		"",
		"\t",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}