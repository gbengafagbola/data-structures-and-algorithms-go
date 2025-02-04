package structural

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World!"

	// Corrected syntax: properly initializing PrinterAdapter
	adapter := PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}

	// Corrected formatting: using proper syntax for string formatting
	returnedMsg := adapter.PrintStored()
	expectedMsg := "Legacy Printer: Hello World!\n"
	if returnedMsg != expectedMsg {
		t.Errorf("Message didn't match: expected %q, got %q", expectedMsg, returnedMsg)
	}

	// Testing when OldPrinter is nil
	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	expectedMsg = "Hello World!"
	if returnedMsg != expectedMsg {
		t.Errorf("Message didn't match: expected %q, got %q", expectedMsg, returnedMsg)
	}
}