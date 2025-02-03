package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return &SportMotorbike{}, nil
	case CruiseMotorbikeType:
		return &CruiseMotorbike{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized", v))
	}
}
