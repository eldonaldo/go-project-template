package user

import (
	"github.com/eldonaldo/go-project-template/core"
)

// Table name for user
func (User) TableName() string {
	return "users"
}

// Package init
func init() {
	// If you create a new model use the commands below once to auto generate
	// SQL statements which are directly issue to the database. The SQL statements
	//can then be used as initial SQL setup to bootstrap your migrations
	
	//userRepository := UserRepository{}
	//_ = userRepository.Db().AutoMigrate(&User{})
}

// user repository
type UserRepository struct {
	core.BaseRepository
}

// Creates a new user
func (receiver UserRepository) CreateNew(name string) *User {
	account := &User{
		Name: name,
	}

	receiver.Db().Create(account)
	return account
}

// Finds the user by name
func (receiver UserRepository) FindByName(name string) *User {
	var user User
	receiver.FindFirstOrPanic(&user, "name = ?", name)
	return &user
}
