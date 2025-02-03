package creational

import "testing"

// TestBuilderPattern tests the Vehicle Builder pattern
func TestBuilderPattern(t *testing.T) {
	// Create a ManufacturingDirector to manage the building process
	manufacturingComplex := ManufacturingDirector{}

	// ---- Test Car Builder ----
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder) // Assign car builder to director
	manufacturingComplex.Construct()           // Build the car
	car := carBuilder.GetVehicle()             // Retrieve the built car

	// Validate car properties
	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4, but got %d", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car', but got %s", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5, but got %d", car.Seats)
	}

	// ---- Test Bike Builder ----
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder) // Assign bike builder to director
	manufacturingComplex.Construct()            // Build the bike
	motorbike := bikeBuilder.GetVehicle()       // Retrieve the built bike

	// Validate bike properties
	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2, but got %d", motorbike.Wheels)
	}

	if motorbike.Structure != "Bike" {
		t.Errorf("Structure on a motorbike must be 'Bike', but got %s", motorbike.Structure)
	}

	if motorbike.Seats != 1 {
		t.Errorf("Seats on a motorbike must be 1, but got %d", motorbike.Seats)
	}

	// ---- Test Bus Builder ----
	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder) // Assign bus builder to director
	manufacturingComplex.Construct()           // Build the bus
	bus := busBuilder.GetVehicle()             // Retrieve the built bus

	// Validate bus properties
	if bus.Wheels != 6 {
		t.Errorf("Wheels on a bus must be 6, but got %d", bus.Wheels)
	}

	if bus.Structure != "Bus" {
		t.Errorf("Structure on a bus must be 'Bus', but got %s", bus.Structure)
	}

	if bus.Seats != 30 {
		t.Errorf("Seats on a bus must be 30, but got %d", bus.Seats)
	}

	// ---- Test Truck Builder ----
	truckBuilder := &TruckBuilder{}
	manufacturingComplex.SetBuilder(truckBuilder) // Assign truck builder to director
	manufacturingComplex.Construct()             // Build the truck
	truck := truckBuilder.GetVehicle()           // Retrieve the built truck

	// Validate truck properties
	if truck.Wheels != 18 {
		t.Errorf("Wheels on a truck must be 18, but got %d", truck.Wheels)
	}

	if truck.Structure != "Truck" {
		t.Errorf("Structure on a truck must be 'Truck', but got %s", truck.Structure)
	}

	if truck.Seats != 2 {
		t.Errorf("Seats on a truck must be 2, but got %d", truck.Seats)
	}
}
