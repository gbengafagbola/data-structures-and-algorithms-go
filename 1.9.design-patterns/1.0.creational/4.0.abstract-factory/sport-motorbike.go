package abstract_factory

type SportMotorbike struct{}

func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}

func (s *SportMotorbike) GetWheels() int {
	return 2
}

func (s *SportMotorbike) GetSeats() int {
	return 1
}
