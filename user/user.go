package user

import (
	"fmt"
	"github.com/eldonaldo/go-project-template/core"
)

// User model
type User struct {
	core.BaseModel
	Name string
}

// Greets the user
func Greet(user *User) string {
	return fmt.Sprintf("Hi %s", user.Name)
}
