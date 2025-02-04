package composition

import "testing"

// TestAthlete_Train ensures that an Athlete can successfully train.
func TestAthlete_Train(t *testing.T) {
	athlete := Athlete{} // Create an Athlete instance
	athlete.Train()      // Call the Train method
}

// TestSwimmer_Swim ensures that CompositeSwimmerA can both train and swim.
func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim // Assign Swim function to a variable

	swimmer := CompositeSwimmerA{
		MyAthlete: Athlete{}, // Embedded Athlete instance
		MySwim:    &localSwim, // Pointer to the Swim function
	}

	swimmer.MyAthlete.Train() // Ensure the athlete can train
	(*swimmer.MySwim)()       // Call the swim function via pointer
}

// TestAnimal_Swim ensures that a Shark can both eat and swim.
func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim, // Assigning the Swim function to Shark
	}

	fish.Eat() // Ensure the shark can eat
	fish.Swim() // Ensure the shark can swim
}

// TestSwimmer_Swim2 ensures that CompositeSwimmerB can train and swim.
func TestSwimmer_Swim2(t *testing.T) {
	swimmer := CompositeSwimmerB{
		Trainer: &Athlete{},            // Implements Trainer interface
		Swimmer: &SwimmerImplementor{}, // Implements Swimmer interface
	}

	swimmer.Train() // Ensure training works
	swimmer.Swim()  // Ensure swimming works
}
