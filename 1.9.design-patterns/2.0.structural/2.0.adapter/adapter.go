package structural

import "fmt"

// LegacyPrinter defines an interface for an old printing system that we need to adapt.
type LegacyPrinter interface {
	Print(s string) string
}

// MyLegacyPrinter is an old printing system that we cannot modify directly.
type MyLegacyPrinter struct{}

// Print prints a message using the old printing system and returns it.
func (l *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg) // Simulating the old behavior of printing directly
	return newMsg
}

// NewPrinter defines an interface for the new printing system.
type NewPrinter interface {
	PrintStored() string
}

// PrinterAdapter adapts the old printer to be compatible with the new system.
type PrinterAdapter struct {
	OldPrinter LegacyPrinter // The old printer being adapted
	Msg        string        // The message to be printed
}

// PrintStored adapts the old printer's method to match the new system's expectations.
func (p *PrinterAdapter) PrintStored() string {
	if p.OldPrinter != nil {
		return p.OldPrinter.Print(p.Msg) // Calls the old printer's Print method
	}
	return p.Msg
}
