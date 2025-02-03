package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType      = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return &CarFactory{}, nil
	case MotorbikeFactoryType:
		return &MotorbikeFactory{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized", f))
	}
}
