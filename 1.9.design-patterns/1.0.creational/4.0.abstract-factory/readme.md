## Abstract Factory Pattern in Go

## Overview
The **Abstract Factory Pattern** (also known as the **Factory of Factories**) is a creational design pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes.

This pattern is useful when multiple factories are needed to create different types of objects that share a common interface.

## Implementation
In this example, we implement an abstract factory for creating different types of **vehicles**:
- **Cars** (LuxuryCar, FamilyCar)
- **Motorbikes** (SportMotorbike, CruiseMotorbike)

Each factory (**CarFactory** and **MotorbikeFactory**) produces specific types of vehicles. We use a top-level **VehicleFactory** that returns the appropriate factory based on input parameters.

## File Structure
```
abstract_factory/
│── car.go
│── family-car.go
│── luxury-car.go
│── car-factory.go
│── motorbike.go
│── sport-motorbike.go
│── cruise-motorbike.go
│── motorbike-factory.go
│── vehicle.go
│── vehicle-factory.go
│── vehicle-factory_test.go
```

## Factory Classes and Methods

### 1. Vehicle Interface
Defines common methods for vehicles:
```go
package abstract_factory

type Vehicle interface {
    GetWheels() int
    GetSeats() int
}
```

### 2. Car Interface & Implementations
```go
package abstract_factory

type Car interface {
    Vehicle
    GetDoors() int
}
```
**LuxuryCar Implementation:**
```go
package abstract_factory

type LuxuryCar struct{}

func (l *LuxuryCar) GetDoors() int { return 4 }
func (l *LuxuryCar) GetWheels() int { return 4 }
func (l *LuxuryCar) GetSeats() int { return 5 }
```

### 3. Motorbike Interface & Implementations
```go
package abstract_factory

type Motorbike interface {
    Vehicle
    GetType() int
}
```
**SportMotorbike Implementation:**
```go
package abstract_factory

type SportMotorbike struct{}

func (s *SportMotorbike) GetType() int { return 1 }
func (s *SportMotorbike) GetWheels() int { return 2 }
func (s *SportMotorbike) GetSeats() int { return 1 }
```

### 4. CarFactory
Creates either a **LuxuryCar** or **FamilyCar**.
```go
package abstract_factory

import "errors"

const (
    LuxuryCarType = 1
    FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
    switch v {
    case LuxuryCarType:
        return &LuxuryCar{}, nil
    case FamilyCarType:
        return &FamilyCar{}, nil
    default:
        return nil, errors.New("Vehicle type not recognized")
    }
}
```

### 5. MotorbikeFactory
Creates either a **SportMotorbike** or **CruiseMotorbike**.
```go
package abstract_factory

import "errors"

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
    switch v {
    case 1:
        return &SportMotorbike{}, nil
    case 2:
        return &CruiseMotorbike{}, nil
    default:
        return nil, errors.New("Motorbike type not recognized")
    }
}
```

### 6. VehicleFactory
The **Factory of Factories**, responsible for returning the appropriate factory based on input.
```go
package abstract_factory

import "errors"

const (
    CarFactoryType      = 1
    MotorbikeFactoryType = 2
)

type VehicleFactory interface {
    NewVehicle(v int) (Vehicle, error)
}

func GetVehicleFactory(f int) (VehicleFactory, error) {
    switch f {
    case CarFactoryType:
        return &CarFactory{}, nil
    case MotorbikeFactoryType:
        return &MotorbikeFactory{}, nil
    default:
        return nil, errors.New("Factory type not recognized")
    }
}
```

## Example Usage
```go
package main

import (
    "fmt"
    "abstract_factory"
)

func main() {
    carFactory, _ := abstract_factory.GetVehicleFactory(abstract_factory.CarFactoryType)
    luxuryCar, _ := carFactory.NewVehicle(abstract_factory.LuxuryCarType)
    fmt.Printf("Luxury car has %d doors and %d seats.\n", luxuryCar.GetDoors(), luxuryCar.GetSeats())
}
```

## New Example: Computer Factory
Another practical example of **Abstract Factory** is a **Computer Factory**, which creates different types of computers:
- **Laptop**
- **Desktop**

### 1. Computer Interface
```go
package abstract_factory

type Computer interface {
    GetRAM() string
    GetProcessor() string
}
```

### 2. Laptop Implementation
```go
package abstract_factory

type Laptop struct{}

func (l *Laptop) GetRAM() string { return "16GB" }
func (l *Laptop) GetProcessor() string { return "Intel i7" }
```

### 3. Desktop Implementation
```go
package abstract_factory

type Desktop struct{}

func (d *Desktop) GetRAM() string { return "32GB" }
func (d *Desktop) GetProcessor() string { return "AMD Ryzen 9" }
```

### 4. Computer Factory
```go
package abstract_factory

import "errors"

type ComputerFactory struct{}

const (
    LaptopType  = 1
    DesktopType = 2
)

func (c *ComputerFactory) NewComputer(cType int) (Computer, error) {
    switch cType {
    case LaptopType:
        return &Laptop{}, nil
    case DesktopType:
        return &Desktop{}, nil
    default:
        return nil, errors.New("Computer type not recognized")
    }
}
```

### 5. Example Usage
```go
package main

import (
    "fmt"
    "abstract_factory"
)

func main() {
    compFactory := abstract_factory.ComputerFactory{}
    laptop, _ := compFactory.NewComputer(abstract_factory.LaptopType)
    fmt.Printf("Laptop specs: %s RAM, %s Processor\n", laptop.GetRAM(), laptop.GetProcessor())
}
```

## Conclusion
The **Abstract Factory Pattern** provides a robust structure for creating object families without specifying their exact classes. This implementation demonstrates how to use it efficiently in Go.

