package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	// Get first instance
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nil")
	}

	// Test first call to AddOne
	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling AddOne for the first time, the count must be 1 but it was %d\n", currentCount)
	}

	// Get second instance
	counter2 := GetInstance()
	if counter2 != counter1 {
		t.Error("Expected same instance in counter2 but got a different instance")
	}

	// Test second call to AddOne
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling AddOne using the second counter, the count must be 2 but was %d\n", currentCount)
	}
}
