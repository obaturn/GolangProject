package model

type ContactUsers struct {
	ContactUsersId uint   `gorm:"primaryKey;autoIncrement"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password"`
	UserId         uint   `gorm:"not null"`
}
