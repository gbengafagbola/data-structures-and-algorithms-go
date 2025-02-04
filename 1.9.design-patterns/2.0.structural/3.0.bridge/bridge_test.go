package structural

import (
	"errors"
	"testing"
)

// TestWriter is a struct used to capture written output in tests.
type TestWriter struct {
	Msg string
}

// Write implements the io.Writer interface for TestWriter.
func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	return 0, errors.New("Content received on Writer was empty") // Fixed return order
}

// TestPrintAPI2 tests the behavior of PrinterImpl2.
func TestPrintAPI2(t *testing.T) {
	// Case 1: PrinterImpl2 without an io.Writer should return an error
	api2 := PrinterImpl2{}
	err := api2.PrintMessage("Hello")
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	expectedErrorMessage := "Not implemented yet" // Adjusted to match PrinterImpl2's implementation
	if err.Error() != expectedErrorMessage {
		t.Errorf("Error message was not correct. \nActual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
	}

	// Case 2: PrinterImpl2 with a valid io.Writer should store the message
	testWriter := TestWriter{} // Fixed struct initialization
	api2 = PrinterImpl2{
		Writer: &testWriter, // Passed TestWriter as io.Writer
	}

	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
