package dto

import "time"

type LoginUserResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
