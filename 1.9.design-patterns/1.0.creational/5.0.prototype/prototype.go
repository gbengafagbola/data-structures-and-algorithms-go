package creational

import (
	"errors"
	"fmt"
)

// ShirtCloner defines the interface for cloning shirts
type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

// Constants representing different shirt colors
const (
	White = 1
	Black = 2
	Blue  = 3
)

// ShirtsCache represents a cache for storing shirt prototypes
type ShirtsCache struct{}

// GetClone returns a cloned instance of a shirt based on the given color
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype // Clone by dereferencing the prototype
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

// ItemInfoGetter defines an interface for retrieving item information
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtColor represents a type for shirt colors
type ShirtColor byte

// Shirt represents a shirt with attributes like Price, SKU, and Color
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo returns the shirt's details
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s', color %d, priced at $%.2f", s.SKU, s.Color, s.Price)
}

// GetShirtsCloner initializes and returns a ShirtsCache instance
func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{} // Returning an instance instead of nil
}

// Predefined shirt prototypes
var (
	whitePrototype = &Shirt{
		Price: 13.00,
		SKU:   "WHITE123",
		Color: White,
	}
	blackPrototype = &Shirt{
		Price: 14.00,
		SKU:   "BLACK123",
		Color: Black,
	}
	bluePrototype = &Shirt{
		Price: 16.00,
		SKU:   "BLUE123",
		Color: Blue,
	}
)

// GetPrice returns the price of the shirt
func (s *Shirt) GetPrice() float32 {
	return s.Price
}
