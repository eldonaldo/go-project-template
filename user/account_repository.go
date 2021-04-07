package user

import (
	"github.com/dchest/uniuri"
	"github.com/eldonaldo/go-project-template/core"
)

// Table name for accounts
func (Account) TableName() string {
	return "accounts"
}

// Package init
func init() {
	// Migrate the schema
	//accountRepository := AccountRepository{}
	//_ = accountRepository.Db().AutoMigrate(&Account{})
}

// Plan repository
type AccountRepository struct {
	core.BaseRepository
}

// Creates a new plan with a unique api key
func (receiver AccountRepository) CreateNew() *Account {
	account := &Account{
		ApiKey: uniuri.NewLen(uniuri.UUIDLen),
		Quota:  20,
	}

	receiver.Db().Create(account)
	return account
}

// Finds the account by api key
func (receiver AccountRepository) FindByApiKey(key string) *Account {
	var account Account
	receiver.FindFirstOrPanic(&account, "api_key = ?", key)
	return &account
}
