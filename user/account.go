package user

import "github.com/eldonaldo/go-project-template/core"

type Account struct {
	core.BaseModel
	ApiKey string
	Quota  int
}
