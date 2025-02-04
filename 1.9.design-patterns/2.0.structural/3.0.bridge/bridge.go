package structural

import (
	"errors"
	"fmt"
	"io"
)

// PrinterAPI is an interface that defines a method for printing messages.
type PrinterAPI interface {
	PrintMessage(string) error
}

// PrinterImpl1 is a concrete implementation of PrinterAPI that prints messages to the console.
type PrinterImpl1 struct{}

// PrintMessage prints a message to the console.
func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg) // Fixed "&s" to "%s" (correct format specifier)
	return nil
}

// PrinterImpl2 is another implementation of PrinterAPI that writes to an io.Writer.
type PrinterImpl2 struct {
	Writer io.Writer
}

// PrintMessage in PrinterImpl2 is not implemented yet and returns an error.
func (p *PrinterImpl2) PrintMessage(msg string) error {
	return errors.New("Not implemented yet")
}

// PrinterAbstraction defines a high-level interface for printing.
type PrinterAbstraction interface {
	Print() error
}

// NormalPrinter is a concrete struct implementing PrinterAbstraction.
type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print calls the PrintMessage method of the PrinterAPI implementation.
func (p *NormalPrinter) Print() error {
	return p.Printer.PrintMessage(p.Msg)
}

// PacktPrinter is another concrete struct implementing PrinterAbstraction.
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print calls PrintMessage and prefixes the message with "Message from packt:"
func (p *PacktPrinter) Print() error {
	return p.Printer.PrintMessage(fmt.Sprintf("Message from packt: %s", p.Msg))
}
