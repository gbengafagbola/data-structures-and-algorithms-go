package main

import (
	"log"
	"os"
	"strings"
	"time"
)

// User struct represents a user with an email and name
type User struct {
	Email string
	Name  string
}

// DataBase is a slice of User representing stored user data
var DataBase = []User{
	{Email: "alexander.davis@example.com", Name: "Alexander Davis"},
	{Email: "alexander.jackson@example.com", Name: "Alexander Jackson"},
	{Email: "avery.williams@example.com", Name: "Avery Williams"},
	{Email: "charlotte.smith@example.com", Name: "Charlotte Smith"},
	{Email: "daniel.miller@example.com", Name: "Daniel Miller"},
	{Email: "ella.smith@example.com", Name: "Ella Smith"},
	{Email: "jacob.white@example.com", Name: "Jacob White"},
	{Email: "james.martinez@example.com", Name: "James Martinez"},
	{Email: "james.miller@example.com", Name: "James Miller"},
	{Email: "jayden.jackson@example.com", Name: "Jayden Jackson"},
	{Email: "liam.robinson@example.com", Name: "Liam Robinson"},
	{Email: "mason.martin@example.com", Name: "Mason Martin"},
	{Email: "matthew.jackson@example.com", Name: "Matthew Jackson"},
	{Email: "mia.smith@example.com", Name: "Mia Smith"},
	{Email: "michael.white@example.com", Name: "Michael White"},
	{Email: "natalie.martin@example.com", Name: "Natalie Martin"},
}

// Worker struct represents a worker that processes a subset of users
type Worker struct {
	users []User  // Slice of users assigned to the worker
	ch    chan *User  // Channel for sending found users
	name  string  // Worker name identifier
}

// NewWorker initializes a new Worker instance
func NewWorker(users []User, ch chan *User, name string) *Worker {
	return &Worker{users: users, ch: ch, name: name}
}

// Find searches for a user email within the worker's assigned users
func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if strings.Contains(user.Email, email) { // Check if email contains search term
			w.ch <- user // Send found user to the channel
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <email>")
	}

	email := os.Args[1] // Get the email argument from command line
	ch := make(chan *User) // Create a channel for user search results

	log.Printf("Looking for %s", email)

	// Split the database among three workers and run them concurrently
	go NewWorker(DataBase[:6], ch, "#1").Find(email)
	go NewWorker(DataBase[6:12], ch, "#2").Find(email)
	go NewWorker(DataBase[12:], ch, "#3").Find(email)

	// Listen for user results or timeout if no result is found
	for {
		select {
		case user := <-ch:
			log.Printf("The email is %s owned by %s", user.Email, user.Name)
		case <-time.After(200 * time.Millisecond):
			log.Printf("The email %s was not found", email)
			return // Exit loop after timeout
		}
	}
}
