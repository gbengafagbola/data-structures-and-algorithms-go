## Go Interfaces Explained with a Railroad and Train Example 🚆

## Overview
This repository demonstrates how **interfaces** work in Go by implementing a simple **railroad width checker**. We define an interface `RailroadWideChecker`, which requires a method to check the width of a train. Then, we use it to determine whether a train can pass through a railroad of a specific width.

## What is an Interface in Go?
An **interface** in Go defines a set of method signatures that a type must implement. Unlike other languages, Go interfaces are **implicit**, meaning a type satisfies an interface simply by implementing its methods.

## How It Works
The example consists of:
- **`RailroadWideChecker` Interface**: Defines a method `CheckRailsWidth()` that returns the width of a train.
- **`Railroad` Struct**: Represents a railroad with a set width.
- **`Train` Struct**: Represents a train with a specific width.
- **Implementation of the Interface**: `Train` implements `RailroadWideChecker` by defining the `CheckRailsWidth()` method.
- **Validation Function**: The `Railroad` struct has a method `IsCorrectSizeTrain()` that checks if a given train's width matches the railroad's width.

## Code Example

```go
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
```

## Explanation
1. **The `RailroadWideChecker` Interface**  
   - It ensures that any type implementing it provides a `CheckRailsWidth()` method returning an `int`.

2. **The `Train` Struct Implements the Interface**  
   - `Train` has a method `CheckRailsWidth()` that returns its width.
   - Since `Train` implements this method, it automatically satisfies the `RailroadWideChecker` interface.

3. **The `Railroad` Struct Uses the Interface**  
   - `Railroad` has a method `IsCorrectSizeTrain(RailroadWideChecker) bool`.
   - It accepts any type that satisfies the `RailroadWideChecker` interface and checks if its width matches the railroad width.

4. **Testing with `main()`**  
   - We create a railroad with a width of `10` units.
   - A passenger train (`width = 10`) and a cargo train (`width = 15`) are created.
   - The function `IsCorrectSizeTrain()` determines whether each train can pass.

## Expected Output
```
Can passenger train pass? true
Can cargo train pass? false
```

## Key Takeaways
✔ **Interfaces provide abstraction**—we can pass any type that implements `CheckRailsWidth()` without knowing its actual implementation.  
✔ **Implicit implementation**—we do not need to explicitly declare that `Train` implements `RailroadWideChecker`.  
✔ **Decoupled design**—we can add more vehicle types (e.g., trams, metros) without modifying existing code, as long as they implement `CheckRailsWidth()`.  

## Running the Code
To run the code, simply execute:
```sh
go run main.go
```