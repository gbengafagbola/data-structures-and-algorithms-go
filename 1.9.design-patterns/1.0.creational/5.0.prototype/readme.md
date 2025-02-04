## Prototype Design Pattern

## Introduction
The **Prototype Design Pattern** is a creational design pattern used when the cost of creating a new object is expensive, and you want to create new objects by copying an existing prototype. Instead of constructing objects manually, a prototype instance is used as a template to create clones.

## Key Concepts
- **Prototype Object**: A predefined instance that serves as a template.
- **Cloning Method**: A method that duplicates the prototype to create new instances.
- **Avoids Costly Instantiations**: Instead of recreating objects from scratch, cloning improves performance and consistency.

---

## Implementation in Go

### Shirt Example
Let's consider a scenario where we need to create different shirt models with predefined attributes. Instead of manually creating new shirts, we use prototypes and clone them.

```go
package main

import (
	"errors"
	"fmt"
)

// Prototype interface
type ItemInfoGetter interface {
	GetInfo() string
}

// Shirt Color Constants
const (
	White = 1
	Black = 2
	Blue  = 3
)

// Shirt struct
type Shirt struct {
	Price float32
	SKU   string
	Color int
}

// GetInfo method to return shirt details
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s', color %d, priced at $%.2f", s.SKU, s.Color, s.Price)
}

// Prototype store
var (
	whitePrototype = &Shirt{Price: 13.00, SKU: "WHITE123", Color: White}
	blackPrototype = &Shirt{Price: 14.00, SKU: "BLACK123", Color: Black}
	bluePrototype  = &Shirt{Price: 16.00, SKU: "BLUE123", Color: Blue}
)

// ShirtCloner interface
type ShirtCloner interface {
	GetClone(color int) (ItemInfoGetter, error)
}

// ShirtsCache struct
type ShirtsCache struct{}

// GetClone method clones a shirt prototype
func (s *ShirtsCache) GetClone(color int) (ItemInfoGetter, error) {
	switch color {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

// GetShirtsCloner initializes the prototype cache
func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}

func main() {
	cloner := GetShirtsCloner()
	shirt, _ := cloner.GetClone(White)
	fmt.Println(shirt.GetInfo())
}
```

### Explanation:
- We define a **Shirt** struct with attributes like `Price`, `SKU`, and `Color`.
- We maintain a **prototype cache** with predefined shirt models.
- We create a **ShirtsCache** struct that implements the cloning method.
- Using the `GetClone` method, we clone an existing prototype instead of instantiating a new object manually.

---

## Another Example: Document Cloning
The prototype pattern is also useful in document management systems where templates need to be copied rather than recreated.

```go
package main

import "fmt"

// Document Prototype Interface
type Document interface {
	Clone() Document
	Print()
}

// Concrete Document Type
type Report struct {
	Title   string
	Content string
}

// Clone method for Report
func (r *Report) Clone() Document {
	return &Report{
		Title:   r.Title,
		Content: r.Content,
	}
}

// Print method for Report
func (r *Report) Print() {
	fmt.Printf("Report Title: %s\nContent: %s\n", r.Title, r.Content)
}

func main() {
	originalReport := &Report{Title: "Monthly Report", Content: "Sales Data..."}
	clonedReport := originalReport.Clone()
	
	clonedReport.Print()
}
```

### Explanation:
- We define a **Document** interface with a `Clone` method.
- The `Report` struct implements the `Clone` method, allowing duplication of reports.
- Instead of manually creating a new report, we simply clone an existing one.

---

## When to Use the Prototype Pattern
- When object creation is **expensive** or **complex**.
- When there is a need to **avoid duplication** of object initialization.
- When creating objects that have **similar properties**.
- When working with **frameworks or libraries** that require cloning capabilities.

---

## Advantages
✅ **Reduces Object Creation Cost** – Useful when instantiation is expensive.
✅ **Encapsulation of Cloning Logic** – The cloning mechanism is encapsulated inside the object.
✅ **Simplifies Object Creation** – Reduces dependencies on object constructors.

## Disadvantages
❌ **Deep vs. Shallow Copy Complexity** – May require deep copying if objects contain pointers.
❌ **Overhead of Maintaining Prototypes** – Requires maintaining prototype objects.

---

## Conclusion
The Prototype Pattern is a powerful creational pattern used to clone existing objects instead of creating new ones from scratch. It is particularly useful in scenarios where object creation is expensive or complex. Whether you’re designing a shirt store, document management system, or even game characters, the prototype pattern can help improve efficiency and maintainability.

---

### References:
- **Design Patterns: Elements of Reusable Object-Oriented Software** by Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides (Gang of Four)
- GoF Design Patterns in Golang

