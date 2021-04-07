package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Sample unit test the user module
func TestUser_SampleTest(t *testing.T) {
	var user = User{Name: "John"}
	assert.Equal(t, user.TableName(), "users")
	assert.Equal(t, Greet(&user), "Hi John")
}