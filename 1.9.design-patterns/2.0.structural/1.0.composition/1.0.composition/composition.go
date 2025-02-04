package structural

// Athlete struct with a Train method
type Athlete struct{}

// Train prints a training message
func (a *Athlete) Train() {
	println("Training")
}

// Swim function (not part of a struct) prints a swimming message
func Swim() {
	println("Swimming!")
}

// CompositeSwimmerA embeds an Athlete and has a pointer to a Swim function
type CompositeSwimmerA struct {
	MyAthlete Athlete // Athlete capabilities (training)
	MySwim    *func() // Pointer to a swim function
}

// -------------------------

// Trainer interface represents an entity that can train
type Trainer interface {
	Train()
}

// Swimmer interface represents an entity that can swim
type Swimmer interface {
	Swim()
}

// SwimmerImplementor provides a concrete implementation of the Swimmer interface
type SwimmerImplementor struct{}

// Swim prints a swimming message
func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

// CompositeSwimmerB embeds Trainer and Swimmer interfaces
// This allows it to have both training and swimming capabilities
type CompositeSwimmerB struct {
	Trainer // Embedded interface for training
	Swimmer // Embedded interface for swimming
}

// -------------------------

// Animal struct with a basic Eat method
type Animal struct{}

// Eat prints an eating message
func (r *Animal) Eat() {
	println("Eating")
}

// Shark struct embeds Animal and has a Swim function
type Shark struct {
	Animal     // Inherits Eat() method
	Swim func() // Function pointer for swimming behavior
}
