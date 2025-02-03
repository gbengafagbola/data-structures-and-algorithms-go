package creational

import (
	"errors"
	"fmt"
)

// PaymentMethod interface defines the method `Pay` that must be implemented by all payment methods.
type PaymentMethod interface {
	Pay(amount float32) string // Pay processes a payment of a given amount and returns a string description.
}

const (
	Cash      = 1 // Constant representing Cash payment method.
	DebitCard = 2 // Constant representing Debit Card payment method.
)

// CashPM represents the Cash payment method.
type CashPM struct{}

// DebitCardPM represents the Debit Card payment method.
type DebitCardPM struct{}

// Pay method for CashPM. Implements the PaymentMethod interface.
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("Paid %.2f using Cash", amount) // Returns a string with the amount paid via Cash.
}

// Pay method for DebitCardPM. Implements the PaymentMethod interface.
func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("Paid %.2f using Debit Card", amount) // Returns a string with the amount paid via Debit Card.
}

// GetPaymentMethod returns a pointer to a PaymentMethod object based on the provided method type.
// If the method is not recognized, an error is returned.
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPM{}, nil // Return a pointer to CashPM if the method is Cash.
	case DebitCard:
		return &DebitCardPM{}, nil // Return a pointer to DebitCardPM if the method is DebitCard.
	default:
		return nil, errors.New("payment method not implemented") // Return an error if the method is not recognized.
	}
}
