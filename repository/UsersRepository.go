package repository

import (
	"Contact/model"
)

type UserRepository interface {
	FindByPhoneNumber(PhoneNumber string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}
