package model

type User struct {
	Id           uint           `gorm:"primaryKey;autoIncrement"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	PhoneNumber  string         `json:"phone_number"`
	Email        string         `json:"email"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	ContactUsers []ContactUsers `gorm:"foreignKey:UserId"`
}
