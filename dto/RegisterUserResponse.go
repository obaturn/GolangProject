package dto

import "time"

type RegisterUserResponse struct {
	Message       string `json:"message"`
	time.Location `json:"time_zone"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}
