package repository

import (
	"Contact/model"
)

type ContactUsersRepository interface {
	FindByPhoneNumber(phoneNumber string) (*model.ContactUsers, error)
	FindByEmail(email string) (*model.ContactUsers, error)
}
