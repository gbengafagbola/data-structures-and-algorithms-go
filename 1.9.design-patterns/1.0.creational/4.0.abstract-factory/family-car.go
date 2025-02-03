package abstract_factory

type FamilyCar struct{}

func (f *FamilyCar) GetDoors() int {
	return 5
}

func (f *FamilyCar) GetWheels() int {
	return 4
}

func (f *FamilyCar) GetSeats() int {
	return 5
}
