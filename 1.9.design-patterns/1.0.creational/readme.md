## Creational Design Patterns: A Final Summary

## Introduction
Creational design patterns focus on how objects are created, ensuring flexibility and efficiency in object instantiation. There are **five** main creational design patterns:

1. **Singleton**
2. **Factory Method**
3. **Abstract Factory**
4. **Builder**
5. **Prototype**

This document provides a simple explanation of each pattern, their use cases, and how they differ from one another.

---

## 1. Singleton Pattern
### Explanation:
The Singleton Pattern ensures that a **class has only one instance** and provides a global access point to it.

### Use Case:
Useful when a **single instance** of a class is required, such as logging, database connections, or configuration settings.

### Example (Database Connection):
```go
package main

import (
	"fmt"
	"sync"
)

type Database struct {}

var instance *Database
var once sync.Once

func GetDatabaseInstance() *Database {
	once.Do(func() {
		instance = &Database{}
	})
	return instance
}

func main() {
	db1 := GetDatabaseInstance()
	db2 := GetDatabaseInstance()
	fmt.Println(db1 == db2) // Output: true (same instance)
}
```

---

## 2. Factory Method Pattern
### Explanation:
The Factory Method Pattern provides an **interface for creating objects** but lets subclasses **decide which class to instantiate**.

### Use Case:
When you need to create **different types of objects** dynamically without specifying the exact class.

### Example (Vehicle Factory):
```go
package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct {}
func (c Car) Drive() string { return "Driving a Car" }

type Bike struct {}
func (b Bike) Drive() string { return "Riding a Bike" }

func GetVehicle(v string) Vehicle {
	if v == "car" {
		return Car{}
	} else if v == "bike" {
		return Bike{}
	}
	return nil
}

func main() {
	vehicle := GetVehicle("car")
	fmt.Println(vehicle.Drive()) // Output: Driving a Car
}
```

---

## 3. Abstract Factory Pattern
### Explanation:
The Abstract Factory Pattern is **a factory of factories** that provides an interface for creating families of related objects.

### Use Case:
When an application needs **multiple families of objects** that must be used together.

### Example (UI Theme Factory):
```go
package main

import "fmt"

type Button interface {
	Render() string
}

type WindowsButton struct {}
func (w WindowsButton) Render() string { return "Rendering Windows Button" }

type MacOSButton struct {}
func (m MacOSButton) Render() string { return "Rendering MacOS Button" }

type UIFactory interface {
	CreateButton() Button
}

type WindowsFactory struct {}
func (w WindowsFactory) CreateButton() Button { return WindowsButton{} }

type MacOSFactory struct {}
func (m MacOSFactory) CreateButton() Button { return MacOSButton{} }

func GetUIFactory(os string) UIFactory {
	if os == "windows" {
		return WindowsFactory{}
	} else if os == "mac" {
		return MacOSFactory{}
	}
	return nil
}

func main() {
	factory := GetUIFactory("mac")
	button := factory.CreateButton()
	fmt.Println(button.Render()) // Output: Rendering MacOS Button
}
```

---

## 4. Builder Pattern
### Explanation:
The Builder Pattern helps in **constructing complex objects step by step**, separating object creation from representation.

### Use Case:
Useful when an object has **many attributes** and should be built in a structured way.

### Example (Building a Pizza):
```go
package main

import "fmt"

type Pizza struct {
	Size   string
	Cheese bool
	Pepperoni bool
}

type PizzaBuilder struct {
	pizza Pizza
}

func (b *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	b.pizza.Size = size
	return b
}

func (b *PizzaBuilder) AddCheese() *PizzaBuilder {
	b.pizza.Cheese = true
	return b
}

func (b *PizzaBuilder) AddPepperoni() *PizzaBuilder {
	b.pizza.Pepperoni = true
	return b
}

func (b *PizzaBuilder) Build() Pizza {
	return b.pizza
}

func main() {
	pizza := new(PizzaBuilder).SetSize("Large").AddCheese().Build()
	fmt.Println(pizza) // Output: {Large true false}
}
```

---

## 5. Prototype Pattern
### Explanation:
The Prototype Pattern allows **copying existing objects** instead of creating new ones from scratch.

### Use Case:
When creating new objects is **costly or complex**, and you want to **clone** existing ones.

### Example (Cloning a Shirt):
```go
package main

import "fmt"

type Shirt struct {
	Color string
}

func (s *Shirt) Clone() *Shirt {
	copy := *s
	return &copy
}

func main() {
	original := &Shirt{Color: "Red"}
	clone := original.Clone()

	fmt.Println(original == clone) // Output: false (Different instances)
	fmt.Println(original.Color, clone.Color) // Output: Red Red
}
```

---

## Differences Between Creational Design Patterns

| Pattern          | Purpose                                     | Example Use Case       |
|-----------------|---------------------------------|------------------|
| **Singleton**   | Ensures only one instance exists | Database Connection |
| **Factory Method** | Provides an interface to create objects dynamically | Vehicle Factory (Car/Bike) |
| **Abstract Factory** | Creates families of related objects | UI Themes (Mac/Windows) |
| **Builder**     | Step-by-step creation of complex objects | Pizza Builder |
| **Prototype**   | Clones an existing object | Shirt Cloning |

---

## Conclusion
Creational design patterns help in managing **object creation efficiently**. Choosing the right pattern depends on the specific problem:
- Use **Singleton** when only **one** instance is needed.
- Use **Factory Method** when you **need to create different types dynamically**.
- Use **Abstract Factory** when creating **families of related objects**.
- Use **Builder** when **constructing complex objects step by step**.
- Use **Prototype** when **cloning objects is more efficient than creating new ones**.

By understanding these patterns, developers can design better, more scalable applications!

---

### References:
- "Design Patterns: Elements of Reusable Object-Oriented Software" by Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides (Gang of Four)
- GoF Design Patterns in Golang

