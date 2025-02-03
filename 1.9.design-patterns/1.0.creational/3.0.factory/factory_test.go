package creational

import (
	"strings"
	"testing"
)

// TestGetPaymentMethodCash tests that the Cash payment method is correctly returned and behaves as expected.
func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "Paid 10.30 using Cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

// TestGetPaymentMethodDebitCard tests that the DebitCard payment method is correctly returned and behaves as expected.
func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "Paid 22.30 using Debit Card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

// TestGetPaymentMethodNonExistent tests that an invalid payment method ID returns an error.
func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
