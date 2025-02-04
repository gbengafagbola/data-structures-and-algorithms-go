# Bridge Pattern (Structural Design Pattern)

## Overview
The **Bridge Pattern** is a structural design pattern that decouples an abstraction from its implementation, allowing them to evolve independently. This pattern is particularly useful when a class has multiple dimensions of variation, such as different ways of implementing an interface.

Instead of using **inheritance**, the **Bridge Pattern** uses **composition**, which improves flexibility and scalability.

## **Key Components of the Bridge Pattern**
1. **Abstraction** – Defines the high-level control layer containing a reference to the implementation.
2. **Refined Abstraction** – Extends the abstraction but uses the same implementation interface.
3. **Implementation (Interface)** – Defines the interface for concrete implementations.
4. **Concrete Implementation** – Implements the interface defined by the implementation.

---

## **Example: Printer API Using the Bridge Pattern**

Let's consider a scenario where we have multiple types of printers, and we need flexibility in defining how messages are printed (e.g., console printing, file writing).

### **Implementation in Go**

```go
package structural

import (
    "errors"
    "fmt"
    "io"
)

// PrinterAPI is the implementation interface
// It allows different printing mechanisms

type PrinterAPI interface {
    PrintMessage(string) error
}

// PrinterImpl1 prints messages to the console

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
    fmt.Println(msg)
    return nil
}

// PrinterImpl2 writes messages to an io.Writer

type PrinterImpl2 struct {
    Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
    if p.Writer == nil {
        return errors.New("You need to pass an io.Writer to PrinterImpl2")
    }
    _, err := fmt.Fprintln(p.Writer, msg)
    return err
}

// PrinterAbstraction defines the abstraction layer

type PrinterAbstraction interface {
    Print() error
}

// NormalPrinter is a concrete abstraction

type NormalPrinter struct {
    Msg     string
    Printer PrinterAPI
}

func (p *NormalPrinter) Print() error {
    return p.Printer.PrintMessage(p.Msg)
}

// PacktPrinter adds extra formatting to the message

type PacktPrinter struct {
    Msg     string
    Printer PrinterAPI
}

func (p *PacktPrinter) Print() error {
    return p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
}
```

### **How It Works**
1. `PrinterAPI` defines the contract for printing.
2. `PrinterImpl1` and `PrinterImpl2` implement different printing methods.
3. `NormalPrinter` and `PacktPrinter` act as abstractions using `PrinterAPI` implementations.
4. The Bridge Pattern allows switching printing mechanisms **without modifying the abstraction**.

---

## **Another Example: Payment System Using Bridge Pattern**

Imagine we are building a payment system where users can pay via **PayPal** or **Credit Card**, and the payments can be processed in different regions (e.g., USA and Europe).

```go
package main

import "fmt"

// PaymentProcessor defines the interface for different payment methods
type PaymentProcessor interface {
    ProcessPayment(amount float64)
}

// PayPalPayment is a concrete implementation of PaymentProcessor
type PayPalPayment struct{}

func (p *PayPalPayment) ProcessPayment(amount float64) {
    fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
}

// CreditCardPayment is another implementation
type CreditCardPayment struct{}

func (c *CreditCardPayment) ProcessPayment(amount float64) {
    fmt.Printf("Processing Credit Card payment of $%.2f\n", amount)
}

// Payment defines the abstraction
type Payment struct {
    Amount     float64
    Processor  PaymentProcessor
}

func (p *Payment) MakePayment() {
    p.Processor.ProcessPayment(p.Amount)
}

func main() {
    // Use PayPal for payments
    paypal := &PayPalPayment{}
    payment1 := Payment{Amount: 100.0, Processor: paypal}
    payment1.MakePayment()

    // Use Credit Card for payments
    creditCard := &CreditCardPayment{}
    payment2 := Payment{Amount: 250.0, Processor: creditCard}
    payment2.MakePayment()
}
```

### **Why Use the Bridge Pattern Here?**
- The **payment methods** (`PayPalPayment`, `CreditCardPayment`) are independent of the **abstraction** (`Payment`).
- We can **add new payment methods** without modifying the `Payment` class.
- This separation allows **greater flexibility and scalability**.

---

## **When to Use the Bridge Pattern?**
Use the Bridge Pattern when:
✔ You want to **decouple abstraction from implementation**.
✔ You expect **multiple variations** of a class and want to avoid deep inheritance trees.
✔ You need **greater flexibility** when adding new implementations.

---

## **Conclusion**
The Bridge Pattern is a powerful way to separate concerns in software design. It improves code maintainability, flexibility, and scalability by ensuring that changes in one part of the system do not affect others.

By implementing the Bridge Pattern, we can manage variations in implementation (such as different payment methods or printing mechanisms) **without modifying the abstraction layer**.

---

## **References**
- [Go Design Patterns](https://golang.design/patterns/)
- [Refactoring Guru - Bridge Pattern](https://refactoring.guru/design-patterns/bridge)

