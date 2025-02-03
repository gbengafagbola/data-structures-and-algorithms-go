

## 🚀 Builder Pattern in Go  

## 📌 Overview  
The **Builder Pattern** is a creational design pattern that provides a **step-by-step approach** to constructing complex objects. It separates the **construction process** from the **final product**, making the code more maintainable and scalable.  

---

## 🏗️ Explanation of the Builder Pattern  

| Component | Role |
|-----------|-----------------------------------------------------------|
| **`BuildProcess` (Interface - Builder)** | Defines steps to build a vehicle. |
| **`CarBuilder` (Concrete Builder)** | Implements `BuildProcess` for cars. |
| **`VehicleProduct` (Product)** | Represents the final built vehicle. |
| **`ManufacturingDirector` (Director)** | Orchestrates the build process. |

---

## 🔑 Key Features of the Pattern  

✅ **Fluent Interface (Method Chaining)**: Each method returns `BuildProcess`, allowing:  
```go
c.SetWheels().SetSeats().SetStructure()
```
to be executed in a single line.  

✅ **Encapsulation**: The construction logic is **separate** from the product structure.  

✅ **Extensibility**: New builders (e.g., `BikeBuilder`, `TruckBuilder`) can be added **without modifying existing code**.  

---

## ⚙️ Implementation in Go  

### 1️⃣ Define the Builder Interface  
The `BuildProcess` interface declares the steps to build a vehicle.  

```go
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}
```

---

### 2️⃣ Implement a Concrete Builder  
The `CarBuilder` struct implements `BuildProcess` to build a car.  

```go
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}
```

---

### 3️⃣ Define the Product  
The `VehicleProduct` struct represents the final object.  

```go
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}
```

---

### 4️⃣ Use a Director to Construct the Object  
The `ManufacturingDirector` orchestrates the building process.  

```go
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
```

---

## ▶️ Usage Example  

```go
package main

import "fmt"

func main() {
	// Create a CarBuilder
	carBuilder := &CarBuilder{}

	// Create a Director and assign the builder
	director := ManufacturingDirector{}
	director.SetBuilder(carBuilder)

	// Construct the vehicle
	director.Construct()

	// Get the final product
	car := carBuilder.GetVehicle()

	// Print the result
	fmt.Printf("Car built with %d wheels, %d seats, and structure: %s\n", car.Wheels, car.Seats, car.Structure)
}
```

---

## ❓ When to Use the Builder Pattern?  

✔️ When you need to create **complex objects** with many optional configurations.  
✔️ When you want to improve **code readability** and **maintainability**.  
✔️ When you need a **step-by-step construction process**.  

---

## 🔄 Extending the Builder  
You can easily add **new builders** for different vehicles (e.g., `BikeBuilder`, `TruckBuilder`) without modifying existing code.  

### Example: **BikeBuilder**  

```go
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
```

---

## 🏆 Conclusion  
The **Builder Pattern** makes object creation **modular, flexible, and maintainable**. It’s particularly useful when constructing **complex objects** with multiple optional attributes.  

✅ **Fluent API** makes chaining methods easier.  
✅ **Encapsulation** ensures the product’s internal structure remains hidden from clients.  
✅ **Scalability** allows easy extension with new builders.  

🚀 Now you have a **well-structured Builder Pattern in Go**! **Happy Coding!** 🎯  

---

This **README** is now **formatted properly** with **clear sections, headings, code formatting, and emojis** for better readability and engagement.