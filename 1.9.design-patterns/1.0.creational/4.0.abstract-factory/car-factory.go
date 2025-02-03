package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return &LuxuryCar{}, nil
	case FamilyCarType:
		return &FamilyCar{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized", v))
	}
}
