package main

import (
	"fmt"
	"strings"
)

// Struct to store weather details of a city
type WeatherData struct {
	Location       string
	MeanTemperature float64 
	Precipitation   float64 
}

// Function to identify the city with the maximum temperature
func getMaxTemperatureCity(data []WeatherData) (string, float64) {
	var warmestCity string
	var maxTemp float64
	for _, record := range data {
		if record.MeanTemperature > maxTemp {
			maxTemp = record.MeanTemperature
			warmestCity = record.Location
		}
	}
	return warmestCity, maxTemp
}

// Function to identify the city with the minimum temperature
func getMinTemperatureCity(data []WeatherData) (string, float64) {
	var coolestCity string
	var minTemp float64 = data[0].MeanTemperature
	for _, record := range data {
		if record.MeanTemperature < minTemp {
			minTemp = record.MeanTemperature
			coolestCity = record.Location
		}
	}
	return coolestCity, minTemp
}

// Function to compute the average precipitation
func calculateAveragePrecipitation(data []WeatherData) float64 {
	var totalPrecipitation float64
	for _, record := range data {
		totalPrecipitation += record.Precipitation
	}
	return totalPrecipitation / float64(len(data))
}

// Function to display cities with precipitation exceeding a threshold
func displayCitiesAboveRainThreshold(data []WeatherData, threshold float64) {
	fmt.Println("\nCities with precipitation greater than", threshold, "mm:")
	for _, record := range data {
		if record.Precipitation > threshold {
			fmt.Printf("%s: %.2f mm\n", record.Location, record.Precipitation)
		}
	}
}

// Function to find and show details of a city based on name
func findCityByName(data []WeatherData, name string) {
	name = strings.ToLower(name)
	found := false
	for _, record := range data {
		if strings.ToLower(record.Location) == name {
			fmt.Printf("\nCity: %s\nMean Temperature: %.2f°C\nPrecipitation: %.2f mm\n", record.Location, record.MeanTemperature, record.Precipitation)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("City not found.")
	}
}

func main() {
	// Weather data for various locations
	weatherRecords := []WeatherData{
		{"New York", 15.2, 120.4},
		{"London", 10.8, 114.5},
		{"Tokyo", 16.3, 120.8},
		{"Paris", 12.5, 90.1},
		{"Sydney", 18.4, 102.3},
	}

	// Determine and display the city with maximum and minimum temperatures
	hottestCity, maxTemp := getMaxTemperatureCity(weatherRecords)
	coolestCity, minTemp := getMinTemperatureCity(weatherRecords)
	fmt.Printf("\nHottest city: %s (%.2f°C)\n", hottestCity, maxTemp)
	fmt.Printf("Coolest city: %s (%.2f°C)\n", coolestCity, minTemp)

	// Compute and show the average precipitation
	averagePrecipitation := calculateAveragePrecipitation(weatherRecords)
	fmt.Printf("\nAverage precipitation: %.2f mm\n", averagePrecipitation)

	// Prompt user for a precipitation threshold and filter cities
	var rainfallThreshold float64
	fmt.Print("\nEnter precipitation threshold (in mm): ")
	_, err := fmt.Scan(&rainfallThreshold)
	if err != nil || rainfallThreshold < 0 {
		fmt.Println("Invalid input. Enter a positive value for the threshold.")
		return
	}
	displayCitiesAboveRainThreshold(weatherRecords, rainfallThreshold)

	// Prompt user for a city name to search
	var searchCity string
	fmt.Print("\nEnter the city name to search: ")
	fmt.Scan(&searchCity)
	findCityByName(weatherRecords, searchCity)
}
