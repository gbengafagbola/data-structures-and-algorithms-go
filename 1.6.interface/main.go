package main

import "fmt"

// Define an interface that requires a method to check rail width
type RailroadWideChecker interface {
	CheckRailsWidth() int
}

// Define a struct representing a railroad with a specific width
type Railroad struct {
	Width int
}

// Method to check if a train's width matches the railroad's width
func (r *Railroad) IsCorrectSizeTrain(p RailroadWideChecker) bool {
	return p.CheckRailsWidth() == r.Width // Returns true if the train width matches the railroad width
}

// Define a struct representing a train with a specific width
type Train struct {
	TrainWidth int
}

// Implement the RailroadWideChecker interface for Train
func (p *Train) CheckRailsWidth() int {
	return p.TrainWidth // Returns the train's width
}

func main() {
	// Create a railroad with a track width of 10 units
	railroad := Railroad{Width: 10}

	// Create two trains: one matching the railroad width, another larger
	passengerTrain := Train{TrainWidth: 10}
	cargoTrain := Train{TrainWidth: 15}

	// Check if each train can pass (must match the railroad width)
	canPassengerTrainPass := railroad.IsCorrectSizeTrain(&passengerTrain) // Should return true
	canCargoTrainPass := railroad.IsCorrectSizeTrain(&cargoTrain)         // Should return false

	// Print the results using %t to correctly format boolean values
	fmt.Printf("Can passenger train pass? %t\n", canPassengerTrainPass)
	fmt.Printf("Can cargo train pass? %t\n", canCargoTrainPass)
}
