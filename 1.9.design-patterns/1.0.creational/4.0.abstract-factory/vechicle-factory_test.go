package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	// Retrieve the motorbike factory
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get a sport motorbike from the factory
	motorbikeVehicle, err := motorbikeF.NewVehicle(SportMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels and %d seats", motorbikeVehicle.GetWheels(), motorbikeVehicle.GetSeats())

	// Assert that the returned vehicle is of type Motorbike
	sportBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sport motorbike has type %d", sportBike.GetType())
}

func TestCarFactory(t *testing.T) {
	// Attempt to retrieve a car factory with an invalid ID
	_, err := GetVehicleFactory(3)
	if err == nil {
		t.Fatal("Car factory with ID 3 should not be recognized")
	}

	// Retrieve the car factory using the correct ID
	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	// Get a luxury car from the factory
	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats and %d wheels", carVehicle.GetSeats(), carVehicle.GetWheels())

	// Assert that the returned vehicle is of type Car
	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has %d doors.", luxuryCar.GetDoors())
}
