package creational

// BuildProcess defines the step-by-step process for building a vehicle
// This is the "Builder" interface in the Builder Pattern
type BuildProcess interface {
	SetWheels() BuildProcess   // Step 1: Set the number of wheels
	SetSeats() BuildProcess    // Step 2: Set the number of seats
	SetStructure() BuildProcess // Step 3: Set the structure (Car, Bike, Truck, Bus)
	GetVehicle() VehicleProduct // Final step: Return the fully built vehicle
}

// ManufacturingDirector orchestrates the building process
// This is the "Director" in the Builder Pattern, ensuring the steps are executed in order
type ManufacturingDirector struct {
	builder BuildProcess // Holds a reference to the builder
}

// Construct directs the builder to assemble a vehicle in a defined sequence
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels() // Calls builder methods in order
}

// SetBuilder assigns a builder to the director
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// VehicleProduct represents the final product that will be built
// This is the "Product" in the Builder Pattern
type VehicleProduct struct {
	Wheels    int    // Number of wheels (e.g., 4 for a car, 2 for a bike)
	Seats     int    // Number of seats
	Structure string // Type of structure (e.g., "Car", "Bike", "Truck", "Bus")
}

// =================== Car Builder ===================

// CarBuilder is a concrete builder that constructs a car
type CarBuilder struct {
	v VehicleProduct // Holds the car being built
}

// SetWheels sets the number of wheels for the car
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats sets the number of seats for the car
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure sets the structure type to "Car"
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle returns the fully built car
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// =================== Bike Builder ===================

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// =================== Bus Builder ===================

type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 6 // Buses typically have 6 wheels
	return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30 // A standard bus has around 30 seats
	return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// =================== Truck Builder ===================

type TruckBuilder struct {
	v VehicleProduct
}

func (t *TruckBuilder) SetWheels() BuildProcess {
	t.v.Wheels = 18 // Trucks typically have 18 wheels
	return t
}

func (t *TruckBuilder) SetSeats() BuildProcess {
	t.v.Seats = 2 // A truck usually has 2 seats
	return t
}

func (t *TruckBuilder) SetStructure() BuildProcess {
	t.v.Structure = "Truck"
	return t
}

func (t *TruckBuilder) GetVehicle() VehicleProduct {
	return t.v
}
