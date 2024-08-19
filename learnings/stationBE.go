package main

import (
	"encoding/json"
	"fmt"
	
)

type StationDetails struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	// Add other details as needed
}

type StationPricing struct {
	ID          int     `json:"id"`
	PricePerKWh float64 `json:"price_per_kwh"`
	// Add other pricing details as needed
}

type CombinedStationData struct {
	Details StationDetails `json:"details"`
	Pricing StationPricing `json:"pricing"`
}

func fetchStationDetails(stationID int) (StationDetails, error) {
	// Your logic to fetch station details goes here
	// ...

	// For illustration, using mock data
	details := StationDetails{
		ID:       stationID,
		Name:     "Station A",
		Location: "City X",
	}
	return details, nil
}

func fetchStationPricing(stationID int) (StationPricing, error) {
	// Your logic to fetch station pricing goes here
	// ...

	// For illustration, using mock data
	pricing := StationPricing{
		ID:          stationID,
		PricePerKWh: 0.12,
	}
	return pricing, nil
}

func combineStationData(stationID int) (CombinedStationData, error) {
	details, err := fetchStationDetails(stationID)
	if err != nil {
		return CombinedStationData{}, err
	}

	pricing, err := fetchStationPricing(stationID)
	if err != nil {
		return CombinedStationData{}, err
	}

	combinedData := CombinedStationData{
		Details: details,
		Pricing: pricing,
	}

	return combinedData, nil
}

func main() {
	stationID := 123
	combinedData, err := combineStationData(stationID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Marshal the combined data to JSON for response or further processing
	resultJSON, _ := json.MarshalIndent(combinedData, "", "  ")
	fmt.Println("Combined Station Data:")
	fmt.Println(string(resultJSON))
}
