package singleton

// singleton struct represents a shared instance that maintains a count.
// This struct should have only one instance throughout the application's lifecycle.
type singleton struct {
	count int
}

// instance is a package-level variable that will hold the single instance of `singleton`.
var instance *singleton

// GetInstance ensures that only one instance of `singleton` is created and returned.
// This follows the Singleton design pattern, which restricts the instantiation of a class to a single object.
func GetInstance() *singleton {
	if instance == nil { // Check if the instance is already created
		instance = new(singleton) // Create a new instance if it doesn't exist
	}
	return instance // Return the single instance
}

// AddOne increments the count variable in the singleton instance and returns the updated value.
// This method helps maintain a global counter that can be accessed consistently across the application.
func (s *singleton) AddOne() int {
	s.count++ // Increase count by 1
	return s.count // Return the updated count value
}
