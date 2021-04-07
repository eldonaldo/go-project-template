package user

import (
	"fmt"
	"github.com/eldonaldo/go-project-template/core"
	"net/http"
)

// Handles user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	name := core.GetRequiredArgument("name", r)

	// Creates a new user
	userRepository := UserRepository{}
	_ = userRepository.CreateNew(name)

	_, _ = fmt.Fprintf(w, "User with name %s created", name)
}

// Handles user greetings
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := core.GetRequiredArgument("name", r)

	// Find user by name
	userRepository := UserRepository{}
	user := userRepository.FindByName(name)

	// And greets him
	_, _ = fmt.Fprintln(w, Greet(user))

	// And also return the user model
	core.PrettyPrint(w, user)
}
